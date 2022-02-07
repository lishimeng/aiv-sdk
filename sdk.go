package aiv_sdk

import (
	"bytes"
	"github.com/lishimeng/aiv-sdk/communication"
	"github.com/lishimeng/aiv-sdk/util"
)

var (
	defaultPacketHeader = PacketHeader{
		First:  PACKET_SYNC_1,
		Second: PACKET_SYNC_2,
	}
)

type Sdk interface {
}

type PacketSdk struct {
	conn      *communication.Conn
	rxInit    bool
	workState int

	handleReply func(ReplyBody)
	handleNote  func(MessageNote)
	handleImage func([]byte)

	enrollQueue bytes.Buffer
}

func (ps *PacketSdk) init() {
	ps.enrollQueue = bytes.Buffer{}
}

func (ps PacketSdk) Builder() (pb *PacketBuilder) {
	pb = &PacketBuilder{}
	return
}

func (ps PacketSdk) send(p PacketTx) (err error) {
	_, err = ps.conn.Write(p.Marshall())
	return
}

func (ps *PacketSdk) StartRxLoop() {
	if ps.rxInit {
		return
	}
	ps.rxInit = true
	ps.conn.Listener = ps.onRx
	go ps.conn.Loop()
}

func (ps *PacketSdk) onRx(buffer bytes.Buffer) {
	// 名称   SyncWord    MsgID    Size    Data    ParityCheck
	// 字节数  2 bytes    1 byte   2 bytes  N bytes 1 byte
	var err error
	var bufferSize int
	var packetSize int16
	var dataSize int16
	for {
		bufferSize = buffer.Len()
		if bufferSize < 2 {
			break
		}
		var bs = buffer.Bytes()
		if bs[0] != byte(PACKET_SYNC_1) { // 保证index[0]正确
			_, err = buffer.ReadByte()
			if err != nil {
				// TODO
			}
			continue
		}
		if bs[1] != byte(PACKET_SYNC_2) { // 保证index[1]正确
			_, err = buffer.ReadByte()
			if err != nil {
				// TODO
			}
			continue
		}

		if bufferSize >= 5 { // 已经取到size字段
			dataSize = util.Bytes2Int16(bs[3:])
			packetSize = dataSize + 2 + 1 + 2 + 1 //sync(2byte)+msgId(1byte)+size(2byte)+chk(1byte)
		}

		if int16(bufferSize) < packetSize {
			return
		}

		var chk = util.CalcChk(bs, packetSize)
		// calc chk

		if chk {
			var messageType = MessageType(bs[2])
			ps.onPacket(messageType, bs[5:], packetSize)
		}

		// clear data
		// move packet
		var i = int16(0)
		for i = 0; i < packetSize; i++ {
			_, _ = buffer.ReadByte()
		}
	}
}

func (ps *PacketSdk) onPacket(messageType MessageType, data []byte, size int16) {
	var rx PacketRx
	rx.MessageType = messageType
	rx.PacketSize = size
	rx.PacketData = data
	switch messageType {
	case MsgReply:
		ps.onReply(rx)
	case MsgNote:
		ps.onNote(rx)
	case MsgImage:
		ps.onImage(rx)
	default:
		// drop
	}
}

func (ps *PacketSdk) onReply(rx PacketRx) {

	var replyBody ReplyBody

	replyBody.Mid = rx.PacketData[0]
	replyBody.Result = rx.PacketData[1]
	replyBody.Data = append(replyBody.Data, rx.PacketData[2:rx.PacketSize+2]...)
	if ps.handleReply != nil {
		go ps.handleReply(replyBody)
	}
}

func (ps *PacketSdk) onNote(rx PacketRx) {
	var note MessageNote

	note.Nid = rx.PacketData[0]
	note.Data = append(note.Data, rx.PacketData[1:rx.PacketSize+1]...)
	if ps.handleNote != nil {
		go ps.handleNote(note)
	}
}

func (ps *PacketSdk) onImage(rx PacketRx) {

	var data []byte
	data = append(data, rx.PacketData[1:rx.PacketSize+1]...)

	if ps.handleImage != nil {
		go ps.handleImage(data)
	}
}