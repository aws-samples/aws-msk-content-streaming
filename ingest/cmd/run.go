package cmd

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Shopify/sarama"
	owm "github.com/briandowns/openweathermap"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

type root struct {
	logger *log.Entry
}

func runE(cmd *cobra.Command, args []string) error {
	// create a new root
	// root := new(root)

	// create root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, gtx := errgroup.WithContext(ctx)

	// data channel
	data := make(chan owm.CurrentWeatherData, 100)

	// schedule weather
	g.Go(fetchWeather(gtx, data))
	g.Go(publishWeather(gtx, data))

	// wait
	if err := g.Wait(); err != nil {
		return err
	}

	// noop
	return nil
}

func publishWeather(ctx context.Context, data <-chan owm.CurrentWeatherData) func() error {
	return func() error {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
		if err != nil {
			return err
		}

		defer func() {
			if err := producer.Close(); err != nil {
				log.Fatalln(err)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return nil
			case w := <-data:
				b, err := json.Marshal(w)
				if err != nil {
					return err
				}

				msg := &sarama.ProducerMessage{Topic: "example.weather", Value: sarama.ByteEncoder(b)}
				partition, offset, err := producer.SendMessage(msg)

				log.Printf("send to partition %s with offset %s", partition, offset)
			}
		}
	}
}

func fetchWeather(ctx context.Context, data chan<- owm.CurrentWeatherData) func() error {
	return func() error {
		ticker := time.NewTicker(10 * time.Second)

		w, err := owm.NewCurrent("C", "en", viper.GetString("api-key")) // fahrenheit (imperial) with Russian output
		if err != nil {
			return err
		}

		for {
			select {
			case <-ticker.C:
				if err := w.CurrentByName("Berlin"); err != nil {
					return err
				}

				data <- *w // we dereference here
			case <-ctx.Done():
				return nil
			}
		}
	}
}
