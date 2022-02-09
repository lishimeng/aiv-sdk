package device

import (
	"bytes"
	"github.com/lishimeng/aiv-sdk/communication"
	"github.com/lishimeng/aiv-sdk/model"
	"github.com/lishimeng/aiv-sdk/util"
)

var (
	DefaultPacketHeader = model.PacketHeader{
		First:  model.PacketSync1,
		Second: model.PacketSync2,
	}
)

type PacketSdk struct {
	conn      *communication.Conn
	rxInit    bool
	workState int

	handleReply func(bool, bool, model.ReplyBody)
	handleNote  func(model.MessageNote)
	handleImage func([]byte)

	enrollQueue bytes.Buffer

	cmdCheck bool
	cmd      model.Command
}

func (ps *PacketSdk) CheckReplyCommand(enable bool) {
	ps.cmdCheck = enable
}

func (ps *PacketSdk) setCommand(cmd model.Command) {
	ps.cmd = cmd
}

func (ps *PacketSdk) SetReplyHandler(handler func(midMatch bool, success bool, body model.ReplyBody)) {
	ps.handleReply = handler
}

func (ps *PacketSdk) init() {
	ps.cmd = model.CmdReady // 应收到ready信号
	ps.enrollQueue = bytes.Buffer{}
}

func (ps PacketSdk) Builder() (pb *PacketBuilder) {
	pb = &PacketBuilder{}
	return
}

func (ps *PacketSdk) send(p model.PacketTx) (err error) {
	cmd := p.Command
	if cmd == model.CmdReset {
		cmd = model.CmdReady // reset 后 收到ready信号
	}
	ps.setCommand(cmd) // 保存当前command
	_, err = ps.conn.Write(p.Marshall())
	return
}

func (ps *PacketSdk) checkCmd(cmd model.Command) bool {
	return ps.cmd == cmd
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
	var packetBodySize int16
	var dataSize int16
	for {
		bufferSize = buffer.Len()
		if bufferSize < 2 {
			break
		}
		var bs = buffer.Bytes()
		if bs[0] != byte(model.PacketSync1) { // 保证index[0]正确
			_, err = buffer.ReadByte()
			if err != nil {
				// TODO
			}
			continue
		}
		if bs[1] != byte(model.PacketSync2) { // 保证index[1]正确
			_, err = buffer.ReadByte()
			if err != nil {
				// TODO
			}
			continue
		}

		if bufferSize >= 5 { // 已经取到size字段
			dataSize = util.Bytes2Int16(bs[3:])
			packetBodySize = dataSize + 1 + 2     // msgId(1byte)+size(2byte)
			packetSize = dataSize + 2 + 1 + 2 + 1 //sync(2byte)+msgId(1byte)+size(2byte)+chk(1byte)
		}

		if int16(bufferSize) < packetSize {
			return
		}

		var chk = util.PacketChk(bs, packetBodySize, bs[packetSize])
		// calc chk

		if chk {
			var messageType = model.MessageType(bs[2])
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

func (ps *PacketSdk) onPacket(messageType model.MessageType, data []byte, size int16) {
	var rx model.PacketRx
	rx.MessageType = messageType
	rx.PacketSize = size
	rx.PacketData = data
	switch messageType {
	case model.MsgReply:
		ps.onReply(rx)
	case model.MsgNote:
		ps.onNote(rx)
	case model.MsgImage:
		ps.onImage(rx)
	default:
		// drop
	}
}

func (ps *PacketSdk) onReply(rx model.PacketRx) {

	var replyBody model.ReplyBody

	var resultCode = model.ResultCode(rx.PacketData[1])
	replyBody.Mid = rx.PacketData[0]
	replyBody.Result = resultCode.Convert()
	replyBody.Data = append(replyBody.Data, rx.PacketData[2:rx.PacketSize+2]...)
	var midMatch = true
	if ps.cmdCheck {
		midMatch = ps.checkCmd(model.Command(replyBody.Mid))
	}
	var success = replyBody.Result.ResultCode == model.MrSuccess.ResultCode

	if ps.handleReply != nil {
		go ps.handleReply(midMatch, success, replyBody)
	}
}

func (ps *PacketSdk) onNote(rx model.PacketRx) {
	var note model.MessageNote

	note.Nid = rx.PacketData[0]
	note.Data = append(note.Data, rx.PacketData[1:rx.PacketSize+1]...)
	if ps.handleNote != nil {
		go ps.handleNote(note)
	}
}

func (ps *PacketSdk) onImage(rx model.PacketRx) {

	var data []byte
	data = append(data, rx.PacketData[1:rx.PacketSize+1]...)

	if ps.handleImage != nil {
		go ps.handleImage(data)
	}
}
