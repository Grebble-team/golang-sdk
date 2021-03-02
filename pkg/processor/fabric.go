package processor

import (
	"fmt"
	"github.com/grebble-team/golang-sdk/pkg/helpers"
)

type AttributeSchema struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

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

func GetAttributesSchema(processor Processor) (map[string]AttributeSchema, error) {
	attributes, err := processor.MapToAttributeType(map[string]string{})
	if err != nil {
		return nil, err
	}
	result := map[string]AttributeSchema{}

	fields := helpers.DeepFields(attributes)
	for _, field := range fields {
		name := field.Tag.Get("json")
		if len(name) > 0 {
			result[name] = AttributeSchema{
				Type:        field.Type.Name(),
				Description: field.Tag.Get("description"),
			}
		}
	}
	return result, nil
}
