package processor

import (
	"fmt"
	"reflect"
)

type AttributeSchema struct {
	Type string `json:"type"`
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
	val := reflect.ValueOf(attributes)
	for i := 0; i < val.Type().NumField(); i++ {
		name := val.Type().Field(i).Tag.Get("json")
		result[name] = AttributeSchema{
			Type: val.Type().Field(i).Type.Name(),
		}
	}
	return result, nil
}
