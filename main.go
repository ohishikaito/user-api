package main

import (
	"fmt"
	"log"
	"net"

	userqueryservice "github.com/ohishikaito/user-api/query/userQueryService"
	definition "github.com/ohishikaito/user-definition"
	"google.golang.org/grpc"
)

type (
	server struct {
		port string
		rpc  *grpc.Server
	}
	// Server interface
	Server interface {
		Serve() error
	}
)

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}

func NewServer() (Server, error) {
	s := &server{}
	s.registerServices()
	return s, nil
}

func (s *server) Serve() error {
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return err
	}
	return s.rpc.Serve(listenPort)
}

func (s *server) registerServices() {
	userQueryService := userqueryservice.NewServer()

	definition.RegisterUserQueryServiceServer(s.rpc, userQueryService)
}
