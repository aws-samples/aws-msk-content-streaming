package server

import (
	"context"
	"net/http"

	"net/http/pprof"
)

var _ Listener = (*debug)(nil)

type debug struct {
	opts    *DebugOpts
	mux     *http.ServeMux
	handler *http.Server

	routes map[string]http.Handler
}

// DebugOpt ...
type DebugOpt func(*DebugOpts)

// DebugOpts ...
type DebugOpts struct {
	Addr string

	Routes map[string]http.Handler
}

// NewDebugListener ...
func NewDebugListener(opts ...DebugOpt) Listener {
	options := &DebugOpts{}
	options.Routes = make(map[string]http.Handler)

	d := new(debug)
	d.opts = options

	// create the mus
	d.mux = http.NewServeMux()

	configureDebug(d, opts...)
	configureMux(d)

	d.handler = new(http.Server)
	d.handler.Addr = d.opts.Addr
	d.handler.Handler = d.mux

	return d
}

// Start ...
func (d *debug) Start(ctx context.Context, ready func()) func() error {
	return func() error {
		// noop, call to be ready
		ready()

		if err := d.handler.ListenAndServe(); err != nil {
			return err
		}

		return nil
	}
}

// WithStatusAddr is adding this status addr as an option.
func WithStatusAddr(addr string) func(o *DebugOpts) {
	return func(o *DebugOpts) {
		o.Addr = addr
	}
}

// WithPprof ...
func WithPprof() func(o *DebugOpts) {
	return func(o *DebugOpts) {
		o.Routes["/debug/pprof/trace"] = http.HandlerFunc(pprof.Trace)
		o.Routes["/debug/pprof/"] = http.HandlerFunc(pprof.Index)
		o.Routes["/debug/pprof/cmdline"] = http.HandlerFunc(pprof.Cmdline)
		o.Routes["/debug/pprof/profile"] = http.HandlerFunc(pprof.Profile)
		o.Routes["/debug/pprof/symbol"] = http.HandlerFunc(pprof.Symbol)
	}
}

// WithPrometheusHandler is adding this prometheus http handler as an option.
func WithPrometheusHandler(handler http.Handler) func(o *DebugOpts) {
	return func(o *DebugOpts) {
		o.Routes["/debug/metrics"] = handler
	}
}

func configureMux(d *debug) error {
	for route, handler := range d.routes {
		d.mux.Handle(route, handler)
	}

	return nil
}

func configureDebug(d *debug, opts ...DebugOpt) error {
	for _, o := range opts {
		o(d.opts)
	}

	d.routes = d.opts.Routes

	return nil
}
