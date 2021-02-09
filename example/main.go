package example

import (
	"encoding/json"
	"fmt"
	"github.com/grebble-team/golang-sdk/pkg/processor"
	"github.com/grebble-team/golang-sdk/server"
)

type ExampleProcessor struct {
	processor.UnimplementedProcessor
}

type Attributes struct {
	Name string `json:"name"`
}

func (e ExampleProcessor) Name() string {
	return "test-processor"
}

func (e ExampleProcessor) MapToAttributeType(attributes string) (interface{}, error) {
	result := Attributes{}
	err := json.Unmarshal([]byte(attributes), &result)
	return result, err
}

func (e ExampleProcessor) Execute(content string, a interface{}, stream processor.Stream) error {
	attribute := a.(Attributes)
	for _, i := range []int{1, 2, 3, 4} {
		err := stream.Send(&processor.StreamResult{
			Content: fmt.Sprintf("%s %s %s", content, string(i), attribute.Name),
			Attributes: map[string]string{
				"test1": "test2",
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var processors []processor.Processor
	processors = append(processors, ExampleProcessor{})
	server.Start(processors)
}
