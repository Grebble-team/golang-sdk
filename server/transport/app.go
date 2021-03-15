package transport

import (
	"context"
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

		attributeSchemaRes := v1.AttributeSchema{
			Items: make(map[string]*v1.AttributeSchemaItem),
		}
		for key, element := range attributeSchema {
			attributeSchemaRes.Items[key] = &v1.AttributeSchemaItem{
				Type:        element.Type,
				Description: element.Description,
			}
		}
		response.Processors = append(response.Processors, &v1.ProcessorExternalInfo{
			Name:            processor.Name(),
			AttributeSchema: &attributeSchemaRes,
		})
	}

	return response, nil
}
