package processor

type StreamResult struct {
	Content     string
	Attributes  map[string]string
	ContentType string
}

type Stream struct {
	Send func(*StreamResult) error
}

type Processor interface {
	Name() string
	Execute(content string, attributes interface{}, stream Stream) error
	MapToAttributeType(attribute string) (interface{}, error)
}

type UnimplementedProcessor struct {
	Processor
}
