package processor

import (
	"github.com/free5gc/smf/internal/sbi/consumer"
	"github.com/free5gc/smf/pkg/app"
)

const (
	CONTEXT_NOT_FOUND = "CONTEXT_NOT_FOUND"
)

type ProcessorSmf interface {
	app.App
}

type Processor struct {
	ProcessorSmf

	consumer *consumer.Consumer
}

func NewProcessor(smf ProcessorSmf, consumer *consumer.Consumer) (*Processor, error) {
	p := &Processor{
		ProcessorSmf: smf,
		consumer:     consumer,
	}
	return p, nil
}

func (p *Processor) Consumer() *consumer.Consumer {
	return p.consumer
}
