package device

import "github.com/lishimeng/aiv-sdk/model"

// StartFaceEnroll
// 开始登记人脸
func (ps *PacketSdk) StartFaceEnroll(admin bool, userName string) {
	// enroll
	ps.enrollQueue.Reset()
	ps.enrollQueue.WriteByte(byte(model.FaceDirectionMiddle))

	ps.enrollQueue.WriteByte(byte(model.FaceDirectionUp))
	ps.enrollQueue.WriteByte(byte(model.FaceDirectionDown))
	ps.enrollQueue.WriteByte(byte(model.FaceDirectionLeft))
	ps.enrollQueue.WriteByte(byte(model.FaceDirectionRight))
	ps.workState = model.WorkStateEnroll
	ps.handleReply = ps.faceEnrollProcess
	// middle

	d, _ := ps.enrollQueue.ReadByte()
	var direction model.FaceDir = model.FaceDir(d)

	var isAdmin byte = 0x00
	if admin {
		isAdmin = 0x01
	}

	err := ps.Enroll(model.MessageEnRoll{
		Admin:    isAdmin,
		UserName: []byte(userName),
		Timeout:  10,
		FaceDir:  direction,
	})

	if err != nil {
		// TODO
		ps.StopFaceEnroll()
	}
}

func (ps *PacketSdk) faceEnrollProcess(midMatch, success bool, body model.ReplyBody) {

	if body.Result == 0x01 {
		// TODO next direction
	}
}

// StopFaceEnroll
// 停止人脸采集
func (ps *PacketSdk) StopFaceEnroll() {
	ps.workState = model.WorkStateIdle
	ps.handleReply = nil
}
