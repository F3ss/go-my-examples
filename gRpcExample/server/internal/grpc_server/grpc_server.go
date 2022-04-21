package grpcserver

import (
	"context"
	"grpc-server/pkg/api"
)

type GRPCServer struct{}

func (g *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (api.AddResponse, error) {
	return api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
