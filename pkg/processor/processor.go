package processor

type StreamResult struct {
	Content     []byte
	Attributes  map[string]string
	ContentType string
}

type Stream struct {
	Send func(*StreamResult) error
}

type Processor interface {
	Name() string
	Execute(content []byte, attributes interface{}, stream Stream) error
	MapToAttributeType(attribute map[string]string) (interface{}, error)
}

type UnimplementedProcessor struct {
	Processor
}
