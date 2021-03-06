package main

import (
	"fmt"
	"github.com/grebble-team/golang-sdk/pkg/helpers"
	"github.com/grebble-team/golang-sdk/pkg/processor"
	"github.com/grebble-team/golang-sdk/server"
	"time"
)

type ExampleProcessor struct {
	processor.UnimplementedProcessor
}

type BaseAttributes struct {
	Id string `json:"id" description:"Id of base attributes"`
}

type Attributes struct {
	Name string `json:"name" description:"Name of base attributes"` //Name or example processor
	BaseAttributes
}

func (e ExampleProcessor) Name() string {
	return "test-processor"
}

func (e ExampleProcessor) MapToAttributeType(attributes map[string]string) (interface{}, error) {
	result := Attributes{}
	err := helpers.MapAttributesContentType(attributes, &result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (e ExampleProcessor) Execute(content []byte, a interface{}, stream processor.Stream) error {
	attribute := a.(Attributes)
	for _, i := range []int{1, 2, 3, 4} {
		err := stream.Send(&processor.StreamResult{
			Content: []byte(fmt.Sprintf("%s %s %s", content, string(i), attribute.Name)),
			Attributes: map[string]string{
				"test1": "test2",
			},
		})
		if err != nil {
			return err
		}
	}

	isBreak := false

	go func() {
		for 1 > 0 {
			time.Sleep(500)
			select {
			case <-stream.Context.Done():
				isBreak = true
				return
			default:
			}
		}
	}()
	return nil
}

func main() {
	var processors []processor.Processor
	processors = append(processors, ExampleProcessor{})
	server.Start(processors)
}
