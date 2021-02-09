package processor

import "fmt"

type StreamResult struct {
	Content    string
	Attributes map[string]string
}

type Stream struct {
	Send func(*StreamResult) error
}

type Processor interface {
	Name() string
	Execute(content string, attributes map[string]string, stream Stream) error
}

type ProcessorsFabric struct {
	Processors []Processor
}

func NewProcessorsFabric(processors []Processor) ProcessorsFabric {
	return ProcessorsFabric{
		Processors: processors,
	}
}

func (p ProcessorsFabric) GetProcessor(name string) (*Processor, error) {
	for _, processor := range p.Processors {
		if processor.Name() == name {
			return &processor, nil
		}
	}

	return nil, fmt.Errorf("Proessor not found")
}
