package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/cero-t/rabbitbeat/config"
)

type Rabbitbeat struct {
	Configuration *config.Config
	done          chan struct{}
	period        time.Duration
}

// Creates beater
func New() *Rabbitbeat {
	return &Rabbitbeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Rabbitbeat) Config(b *beat.Beat) error {

	// Load beater configuration
	err := cfgfile.Read(&bt.Configuration, "")
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	return nil
}

func (bt *Rabbitbeat) Setup(b *beat.Beat) error {

	// Setting default period if not set
	if bt.Configuration.Rabbitbeat.Period == "" {
		bt.Configuration.Rabbitbeat.Period = "1s"
	}

	var err error
	bt.period, err = time.ParseDuration(bt.Configuration.Rabbitbeat.Period)
	if err != nil {
		return err
	}

	return nil
}

func (bt *Rabbitbeat) Run(b *beat.Beat) error {
	logp.Info("rabbitbeat is running! Hit CTRL-C to stop it.")

	ticker := time.NewTicker(bt.period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		b.Events.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Rabbitbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Rabbitbeat) Stop() {
	close(bt.done)
}
