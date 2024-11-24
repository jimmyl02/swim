package swim

import (
	"context"
	"net"

	"github.com/jimmyl02/swim/types"
)

// swimGRPCServer defines the gRPC server for SWIM
type swimGRPCServer struct {
	types.UnimplementedSwimNodeServer

	// identifiers
	node string // node tracks the name / identifier of the node
	addr net.Addr
	port uint32

	// internal state for the node
	incarnation uint32

	// internal state for the cluster
	nodes map[string]NodeState

	// tracking events
}

// Ping responds if the node is alive and piggybacks gossip
func (s *swimGRPCServer) Ping(ctx context.Context, req *types.PingRequest) (*types.AckResponse, error) {
	return &types.AckResponse{}, nil
}

func NewSwimServer() {

}
