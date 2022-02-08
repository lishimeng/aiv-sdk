package model

const (
	PacketSync1 PacketSync = 0xEF
	PacketSync2 PacketSync = 0xAA
)

type PacketSync byte

type Command byte

type PacketSize struct {
	High byte
	Low  byte
}

type PacketData []byte

type PacketChk byte

type PacketHeader struct {
	First  PacketSync
	Second PacketSync
}

type PacketTx struct {
	PacketHeader
	Command
	PacketSize
	PacketData
	PacketChk
}

type PacketRx struct {
	MessageType
	PacketSize int16
	PacketData
}

func (p *PacketTx) Build() {
	// 计算size
	// 计算chk
}

func (ps PacketSize) Size() int16 {
	var low = int16(ps.Low & 0xff)
	var high = int16(ps.High&0xff) * 256
	var a = low + high

	return a
}

func (p PacketTx) Convert() (c Packet) {

	return
}

func (p *PacketTx) CalcChk() (chk byte) {
	// TODO 累加计算？
	for _, b := range p.PacketData {
		chk = chk + b
	}
	return
}

func (p PacketTx) Marshall() (bs []byte) {
	return
}

type Packet struct {
	Command
	Size int
	Data []byte
}
