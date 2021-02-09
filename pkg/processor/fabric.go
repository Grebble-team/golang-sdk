package processor

import "fmt"

type ProcessorsFabric struct {
	Processors []Processor
}

func NewProcessorsFabric(processors []Processor) ProcessorsFabric {
	return ProcessorsFabric{
		Processors: processors,
	}
}

func (p ProcessorsFabric) GetProcessor(name string) (Processor, error) {
	for _, processor := range p.Processors {
		if processor.Name() == name {
			return processor, nil
		}
	}

	return nil, fmt.Errorf("Processor not found")
}
