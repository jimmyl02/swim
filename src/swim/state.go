package swim

import (
	"net"
)

// NodeStateType represents the types of node states
type NodeStateType int32

const (
	NodeStateType_alive   NodeStateType = 0
	NodeStateType_suspect NodeStateType = 1
	NodeStateType_dead    NodeStateType = 2
)

// NodeState represents the state of a node
type NodeState struct {
	Name        string
	SourceAddr  net.Addr
	SourcePort  uint32
	Sequence    uint32
	State       NodeStateType
	StateChange int64
}
