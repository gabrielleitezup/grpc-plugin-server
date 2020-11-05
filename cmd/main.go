package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"pluginserver/internal"
	"pluginserver/pkg/datasource"
)

func newServer() *internal.DatasourceServer {
	s := &internal.DatasourceServer{}
	return s
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:9000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	datasource.RegisterDatasourceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
