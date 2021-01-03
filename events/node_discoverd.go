package events

import "github.com/justinkiang/gozwave/serialapi"

type NodeDiscoverd struct {
	Address int

	serialapi.FuncGetNodeProtocolInfo
}
