package transport

import (
	"encoding/json"
	v1 "github.com/grebble-team/golang-sdk/pkg/grpc/inner/v1"
	pkgprocessor "github.com/grebble-team/golang-sdk/pkg/processor"
)

type ProcessorServer struct {
	v1.UnimplementedProcessorServer
	ProcessorsFabric pkgprocessor.ProcessorsFabric
}

func NewProcessorServer(processorsFabric pkgprocessor.ProcessorsFabric) ProcessorServer {
	return ProcessorServer{
		ProcessorsFabric: processorsFabric,
	}
}

func (p ProcessorServer) Execute(req *v1.FlowExecuteRequest, stream v1.Processor_ExecuteServer) error {
	processor, err := p.ProcessorsFabric.GetProcessor(req.FlowName)
	if err != nil {
		return err
	}

	attributes, err := processor.MapToAttributeType(req.Attributes)
	if err != nil {
		return err
	}

	processor.Execute(req.Content, attributes, pkgprocessor.Stream{
		Send: func(req *pkgprocessor.StreamResult) error {
			attr, err := json.Marshal(req.Attributes)
			if err != nil {
				return err
			}
			mimeType := "text/plain"
			if len(req.ContentType) > 0 {
				mimeType = req.ContentType
			}
			return stream.Send(&v1.FlowExecuteResponse{
				Content:    req.Content,
				Attributes: string(attr),
				StreamEnd:  false,
				MimeType:   mimeType,
			})
		},
	})

	return stream.Send(&v1.FlowExecuteResponse{
		Content:    "",
		Attributes: "",
		StreamEnd:  true,
	})

	return nil
}
