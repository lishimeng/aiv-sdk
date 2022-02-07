package aiv_sdk

type PacketBuilder struct {
	Command
	payload PacketData
}

func (pb *PacketBuilder) Cmd(cmd Command) *PacketBuilder {
	pb.Command = cmd
	return pb
}

func (pb *PacketBuilder) Body(payload PacketData) *PacketBuilder {
	pb.payload = payload
	return pb
}

func (pb *PacketBuilder) Build() (p PacketTx) {

	// 设置sync字段2byte
	p.PacketHeader = defaultPacketHeader
	// 设置cmd字段1byte
	p.Command = pb.Command
	// 设置数据字段N byte
	p.PacketData = pb.payload
	p.Build()
	return
}
