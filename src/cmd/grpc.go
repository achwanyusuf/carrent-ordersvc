package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var listener net.Listener

type GRPC struct {
	Host string
	Port int
}

func (g *GRPC) newGRPC() *grpc.Server {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", g.Host, g.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	listener = lis
	return grpc.NewServer()
}

// func (g *GRPC) newGRPCClient() (conn, grpcmodel.OrderClient) {
// 	// Set up a connection to the server.
// 	conn, err := grpc.Dial(fmt.Sprintf("%s:%v", g.Host, g.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	orderClient := grpcmodel.NewOrderClient(conn)

// 	return orderClient
// }
