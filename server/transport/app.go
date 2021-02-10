package transport

import (
	"context"
	"encoding/json"
	v1 "github.com/grebble-team/golang-sdk/pkg/grpc/inner/v1"
	pkgprocessor "github.com/grebble-team/golang-sdk/pkg/processor"
)

type AppServer struct {
	v1.UnimplementedExternalAppServer
	ProcessorsFabric pkgprocessor.ProcessorsFabric
}

func NewAppServer(processorsFabric pkgprocessor.ProcessorsFabric) AppServer {
	return AppServer{
		ProcessorsFabric: processorsFabric,
	}
}

func (p AppServer) AppInfo(ctx context.Context, req *v1.AppExternalInfoRequest) (*v1.AppInfoExternalResponse, error) {

	response := &v1.AppInfoExternalResponse{}
	for _, processor := range p.ProcessorsFabric.Processors {
		attributeSchema, err := pkgprocessor.GetAttributesSchema(processor)
		if err != nil {
			return nil, err
		}
		schemaJson, err := json.Marshal(attributeSchema)
		if err != nil {
			return nil, err
		}
		response.Processors = append(response.Processors, &v1.ProcessorExternalInfo{
			Name:            processor.Name(),
			AttributeSchema: string(schemaJson),
		})
	}

	return response, nil
}
