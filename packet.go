package aiv_sdk

import "fmt"

const (
	PACKET_SYNC_1 PacketSync = 0xEF
	PACKET_SYNC_2 PacketSync = 0xAA
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

func (p PacketTx) CheckSum() bool {
	return false
}

func (ps PacketSize) Size() int32 {
	var low = int32(ps.Low & 0xff)
	var high = int32(ps.High&0xff) * 256
	var a = low + high

	fmt.Printf("%d", a)
	return a
}

func (p PacketTx) Convert() (c Packet) {

	return
}

func (p PacketTx) CalcChk() byte {
	// TODO
	return 0x00
}

func (p PacketTx) Marshall() (bs []byte) {
	return
}

type Packet struct {
	Command
	Size int
	Data []byte
}
