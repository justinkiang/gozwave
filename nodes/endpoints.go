package nodes

import (
	"github.com/justinkiang/gozwave/commands"
	"github.com/justinkiang/gozwave/database"
	"github.com/justinkiang/gozwave/interfaces"
	"github.com/sirupsen/logrus"
)

type Endpoint struct {
	Id             int
	CommandClasses []*database.CommandClass

	StateBool  map[string]bool
	StateFloat map[string]float64

	node *Node
}

func (n *Node) Endpoint(id int) *Endpoint {
	if n == nil {
		logrus.Errorf("Failed to get endpoint from NIL node")
		return nil
	}

	if id > len(n.Endpoints) {
		return nil
	}

	return n.Endpoints[id]
}

func (e *Endpoint) Write(msg interfaces.Encodable) {
	logrus.Debugf("Send to endpoint %d ", e.Id)
	e.node.connection.Write(commands.NewMultiChannelEncap(msg.Encode(), e.Id))
}

func (e *Endpoint) On() {
	var send interfaces.Encodable

	switch {
	case e.node.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(true)
		cmd.SetNode(e.node.Id)

		send = cmd
	case e.node.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(100)
		cmd.SetNode(e.node.Id)

		send = cmd
	default:
		return
	}

	e.Write(send)
}

func (e *Endpoint) Off() {
	var send interfaces.Encodable

	switch {
	case e.node.HasCommand(commands.SwitchBinary):
		cmd := commands.NewSwitchBinary()
		cmd.SetValue(false)
		cmd.SetNode(e.node.Id)

		send = cmd
	case e.node.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(0)
		cmd.SetNode(e.node.Id)

		send = cmd
	default:
		return
	}

	e.Write(send)
}

func (e *Endpoint) Level(value float64) {
	var send interfaces.Encodable

	switch {
	case e.node.HasCommand(commands.SwitchMultilevel):
		cmd := commands.NewSwitchMultilevel()
		cmd.SetValue(value)
		cmd.SetNode(e.node.Id)

		send = cmd
	default:
		return
	}

	e.Write(send)
}
