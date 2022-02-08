package device

import (
	"github.com/lishimeng/aiv-sdk/model"
)

type PacketBuilder struct {
	model.Command
	payload model.PacketData
}

func (pb *PacketBuilder) Cmd(cmd model.Command) *PacketBuilder {
	pb.Command = cmd
	return pb
}

func (pb *PacketBuilder) Body(payload model.PacketData) *PacketBuilder {
	pb.payload = payload
	return pb
}

func (pb *PacketBuilder) Build() (p model.PacketTx) {

	// 设置sync字段2byte
	p.PacketHeader = DefaultPacketHeader
	// 设置cmd字段1byte
	p.Command = pb.Command
	// 设置数据字段N byte
	p.PacketData = pb.payload
	p.Build()
	return
}
