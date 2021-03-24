package server

import (
	"fmt"
	v1 "github.com/grebble-team/golang-sdk/pkg/grpc/inner/v1"
	"github.com/grebble-team/golang-sdk/pkg/processor"
	"github.com/grebble-team/golang-sdk/pkg/setting"
	"github.com/grebble-team/golang-sdk/server/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

type Server struct {
	ProcessorServer transport.ProcessorServer
	AppServer       transport.AppServer
}

func NewServer(processors []processor.Processor) Server {
	processorsFabric := processor.NewProcessorsFabric(processors)
	return Server{
		ProcessorServer: transport.NewProcessorServer(processorsFabric),
		AppServer:       transport.NewAppServer(processorsFabric),
	}
}

func (s Server) Start() error {
	grpcEndPoint := fmt.Sprintf(":%d", setting.Settings.Server.GrpcPort)
	lis, err := net.Listen("tcp", grpcEndPoint)
	if err != nil {
		panic(err)
	}

	//TODO Set as parameters
	maxRecvMsgSize := 1024 * 1024 * 256
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(maxRecvMsgSize), //256Mi
		grpc.MaxSendMsgSize(maxRecvMsgSize), //256Mi
	)
	fmt.Printf("MaxRecvMsgSize server %d", maxRecvMsgSize)
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	fmt.Printf("Start server")

	v1.RegisterProcessorServer(grpcServer, s.ProcessorServer)
	v1.RegisterExternalAppServer(grpcServer, s.AppServer)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}

func Start(processors []processor.Processor) {
	server := NewServer(processors)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
