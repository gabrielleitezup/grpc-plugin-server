package main

import (
	"fmt"
	"github.com/gabrielleitezup/grpc-plugin-server/internal"
	"github.com/gabrielleitezup/grpc-plugin-server/pkg/datasource"
	"google.golang.org/grpc"
	"log"
	"net"
)

func newServer() *internal.DatasourceServer {
	s := &internal.DatasourceServer{}
	return s
}
func newRonaldoServer() *internal.RonaldoServer {
	s := &internal.RonaldoServer{}
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
	datasource.RegisterDatasourceServer(grpcServer, newRonaldoServer())
	grpcServer.Serve(lis)
}
