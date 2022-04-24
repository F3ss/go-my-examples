package serv

import (
	"context"
	"grps-server/pkg/api"
)

type GRPCServer struct{}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}

func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
