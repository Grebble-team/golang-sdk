package main

import (
	"fmt"
	"github.com/grebble-team/golang-sdk/pkg/processor"
	"github.com/grebble-team/golang-sdk/server"
)

type ExampleProcessor struct {
	processor.Processor
}

func (e ExampleProcessor) Name() string {
	return "test-processor"
}

func (e ExampleProcessor) Execute(content string, attributes map[string]string, stream processor.Stream) error {
	for _, i := range []int{1, 2, 3, 4} {
		err := stream.Send(&processor.StreamResult{
			Content: fmt.Sprintf("%s %s", content, string(i)),
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
