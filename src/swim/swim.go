package swim

import (
	"context"

	"github.com/jimmyl02/swim/types"
)

// swimGRPCServer defines the gRPC server for SWIM
type swimGRPCServer struct {
	types.UnimplementedSwimNodeServer
}

// Ping responds if the node is alive and piggybacks gossip
func (s *swimGRPCServer) Ping(ctx context.Context, req *types.PingRequest) (*types.AckResponse, error) {
	return &types.AckResponse{}, nil
}

// SwimServer defines a single node in a SWIM cluster
type SwimServer struct {
}

func NewSwimServer() {

}
