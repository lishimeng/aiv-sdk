package aiv_sdk

// StartFaceEnroll
// 开始登记人脸
func (ps *PacketSdk) StartFaceEnroll(admin bool, userName string) {
	// enroll
	ps.enrollQueue.Reset()
	ps.enrollQueue.WriteByte(byte(FaceDirectionMiddle))

	ps.enrollQueue.WriteByte(byte(FaceDirectionUp))
	ps.enrollQueue.WriteByte(byte(FaceDirectionDown))
	ps.enrollQueue.WriteByte(byte(FaceDirectionLeft))
	ps.enrollQueue.WriteByte(byte(FaceDirectionRight))
	ps.workState = WorkStateEnroll
	ps.handleReply = ps.faceEnrollProcess
	// middle

	d, _ := ps.enrollQueue.ReadByte()
	var direction FaceDir = FaceDir(d)

	var isAdmin byte = 0x00
	if admin {
		isAdmin = 0x01
	}

	err := ps.Enroll(MessageEnRoll{
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

func (ps *PacketSdk) faceEnrollProcess(body ReplyBody) {

	if body.Result == 0x01 {
		// TODO next direction
	}
}

// StopFaceEnroll
// 停止人脸采集
func (ps *PacketSdk) StopFaceEnroll() {
	ps.workState = WorkStateIdle
	ps.handleReply = nil
}
