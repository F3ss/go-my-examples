package main

import (
	"grps-server/pkg/api"
	"grps-server/pkg/serv"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := serv.NewGRPCServer()
	api.RegisterAdderServer(s, srv)

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalln(err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
