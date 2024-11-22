package swim

import "github.com/jimmyl02/swim/types"

// swimGRPCServer defines the gRPC server for SWIM
type swimGRPCServer struct {
	types.UnimplementedSwimServer
}

// SwimServer defines a single node in a SWIM cluster
type SwimServer struct {
}

func NewSwimServer() {

}
