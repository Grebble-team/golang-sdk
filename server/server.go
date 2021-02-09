package server

import (
	"fmt"
	v1 "github.com/grebble-team/golang-sdk/pkg/grpc/inner/v1"
	"github.com/grebble-team/golang-sdk/pkg/processor"
	"github.com/grebble-team/golang-sdk/pkg/setting"
	"github.com/grebble-team/golang-sdk/server/transport"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	ProcessorServer transport.ProcessorServer
}

func NewServer(processors []processor.Processor) Server {
	processorsFabric := processor.NewProcessorsFabric(processors)
	return Server{
		ProcessorServer: transport.NewProcessorServer(processorsFabric),
	}
}

func (s Server) Start() error {
	grpcEndPoint := fmt.Sprintf(":%d", setting.Settings.Server.GrpcPort)
	lis, err := net.Listen("tcp", grpcEndPoint)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(52428800),
		grpc.MaxSendMsgSize(52428800),
	)
	fmt.Printf("Start server")

	v1.RegisterProcessorServer(grpcServer, s.ProcessorServer)
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
