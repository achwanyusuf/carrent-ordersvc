package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"

	"github.com/achwanyusuf/carrent-lib/pkg/logger"
	"github.com/achwanyusuf/carrent-ordersvc/script/cred"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var listener net.Listener

type GRPC struct {
	Host string
	Port int
	Log  logger.Logger
}

func loadTLSCredentials(serverCert string, serverKey string) (credentials.TransportCredentials, error) {
	fmt.Println("kena bang", cred.Path(serverCert))
	cert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, err
	}

	return credentials.NewServerTLSFromCert(&cert), nil
}

func loadClientTLSCredentials(clientCert string, clientHost string) (credentials.TransportCredentials, error) {
	fmt.Println("kena bang", cred.Path(clientCert))
	creds, err := credentials.NewClientTLSFromFile(clientCert, clientHost)
	if err != nil {
		return nil, err
	}
	return creds, nil
}

func (g *GRPC) newGRPC(serverCert string, serverKey string) *grpc.Server {
	tlsCredentials, err := loadTLSCredentials(serverCert, serverKey)
	if err != nil {
		g.Log.Panic(context.Background(), "cannot load TLS credentials:", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%v", g.Host, g.Port))
	if err != nil {
		g.Log.Panic(context.Background(), "failed to listen:", err)
	}
	listener = lis
	return grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
}
