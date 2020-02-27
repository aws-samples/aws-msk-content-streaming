package server

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Server is the interface to be implemented
// to run the server.
//
// 	s, ctx := WithContext(context.Background())
//	s.Listen(listener, false)
//
//	if err := s.Wait(); err != nil {
//		panic(err)
//	}
type Server interface {
	// Run is running a new routine
	Listen(listener Listener, ready bool)
	// Name
	Name() string
	// Env
	Env() string
	// Waits for the server to fail,
	// or gracefully shutdown if called
	Wait() error
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	// Name
	Name string
	// Env
	Env string
	// ReloadSignal
	ReloadSignal syscall.Signal
	// TermSignal
	TermSignal syscall.Signal
	// KillSignal
	KillSignal syscall.Signal
}

type listeners map[Listener]bool

// server holds the instance info of the server
type server struct {
	ctx    context.Context
	cancel context.CancelFunc

	wg      sync.WaitGroup
	errOnce sync.Once
	err     error

	listeners map[Listener]bool

	ready chan bool
	sys   chan os.Signal

	opts *Opts
}

// WithContext ...
func WithContext(ctx context.Context, opts ...Opt) (Server, context.Context) {
	ctx, cancel := context.WithCancel(ctx)

	// new server
	s := newServer(ctx, opts...)
	s.cancel = cancel
	s.ctx = ctx

	return s, ctx
}

// Name ...
func Name(name string) func(o *Opts) {
	return func(o *Opts) {
		o.Name = name
	}
}

// Env ...
func Env(env string) func(o *Opts) {
	return func(o *Opts) {
		o.Env = env
	}
}

func newServer(ctx context.Context, opts ...Opt) *server {
	options := &Opts{}

	s := new(server)
	s.opts = options

	ctx, cancel := context.WithCancel(ctx)
	s.cancel = cancel

	s.listeners = make(listeners)
	s.ready = make(chan bool)

	configure(s, opts...)
	configureSignals(s)

	s.Env()

	return s
}

// Listen ...
func (s *server) Listen(listener Listener, ready bool) {
	// if we found the listener already, we simply reject to be added
	if _, found := s.listeners[listener]; found {
		return
	}

	// add to listeners
	s.listeners[listener] = ready
}

// Env ...
func (s *server) Env() string {
	return s.opts.Env
}

// Name ...
func (s *server) Name() string {
	return s.opts.Name
}

// Wait ...
func (s *server) Wait() error {
	// create ticker for interrupt signals
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

OUTTER:
	// start all listeners
	for l, ready := range s.listeners {
		fn := func() {
			r := ready

			var readyOnce sync.Once
			readyOnce.Do(func() {
				if r {
					s.ready <- true
				}
			})
		}

		// schedule to routines
		s.run(l.Start(s.ctx, fn))

		// this blocks until ready is called
		if ready {
			select {
			case <-s.ready:
				continue
			case <-s.ctx.Done():
				break OUTTER
			}
		}
	}

	// this is the main loop
	for {
		select {
		case <-ticker.C:
		case <-s.sys:
			// if there is sys interupt
			// cancel the context of the routines
			s.cancel()
		case <-s.ctx.Done():
			if err := s.ctx.Err(); err != nil {
				return err
			}

			return nil
		}
	}
}

func (s *server) run(f func() error) {
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()

		if err := f(); err != nil {
			s.errOnce.Do(func() {
				s.err = err
				if s.cancel != nil {
					s.cancel()
				}
			})
		}
	}()
}

// Listener is the interface to a listener,
// so starting and shutdown of a listener,
// or any routine.
type Listener interface {
	// Start is being called on the listener
	Start(ctx context.Context, ready func()) func() error
}

func configureSignals(s *server) {
	s.sys = make(chan os.Signal, 1)

	signal.Notify(s.sys, s.opts.ReloadSignal, s.opts.KillSignal, s.opts.TermSignal)
}

func configure(s *server, opts ...Opt) error {
	for _, o := range opts {
		o(s.opts)
	}

	s.opts.TermSignal = syscall.SIGTERM
	s.opts.ReloadSignal = syscall.SIGHUP
	s.opts.KillSignal = syscall.SIGINT

	return nil
}
