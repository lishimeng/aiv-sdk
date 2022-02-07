package aiv_sdk

func (ps *PacketSdk) GetFeatureInfo(param UserId) (err error) {
	var payload []byte
	payload = append(payload, param.UserIdHeb, param.UserIdLeb)
	packet := ps.Builder().Cmd(CmdGetFeatureInfo).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) UploadFeature(param MessageUploadFaceFeature) (err error) {
	var payload []byte
	payload = append(payload, param.Offset...)
	payload = append(payload, param.Size...)
	packet := ps.Builder().Cmd(CmdUploadFeature).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetLogFile() (err error) {
	packet := ps.Builder().Cmd(CmdGetLogFile).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) UploadLogFile(param MessageUploadLogFileData) (err error) {
	var payload []byte
	payload = append(payload, param.LogFileOffset...)
	payload = append(payload, param.LogFileSize...)
	packet := ps.Builder().Cmd(CmdUploadLogFile).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) PowerDown() (err error) {
	packet := ps.Builder().Cmd(CmdPowerDown).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) DemoMode(enable DemoMode) (err error) {
	var payload []byte
	payload = append(payload, byte(enable))
	packet := ps.Builder().Cmd(CmdDemoMode).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) TransFilePacket(param MessageTransFileData) (err error) {
	var payload []byte
	payload = append(payload, param.StoreType)
	payload = append(payload, param.FileSize...)
	payload = append(payload, param.PacketSize...)
	payload = append(payload, param.Data...)
	packet := ps.Builder().Cmd(CmdTransFilePacket).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) InitEncryption(param MessageInitEncryptionData) (err error) {
	var payload []byte
	payload = append(payload, param.Seed...)
	payload = append(payload, param.Mode)
	payload = append(payload, param.CreateTime...)
	packet := ps.Builder().Cmd(CmdInitEncryption).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetVersion() (err error) {
	packet := ps.Builder().Cmd(CmdGetVersion).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) DeleteAll() (err error) {
	packet := ps.Builder().Cmd(CmdDeleteAll).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) GetUserInfo(param UserId) (err error) {
	var payload []byte
	payload = append(payload, param.UserIdHeb, param.UserIdLeb)
	packet := ps.Builder().Cmd(CmdGetUserInfo).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) FaceReset() (err error) {
	packet := ps.Builder().Cmd(CmdFaceReset).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) UploadImage(param MessageUploadImageData) (err error) {
	var payload []byte
	payload = append(payload, param.UploadImageOffset...)
	payload = append(payload, param.UploadImageSize...)
	packet := ps.Builder().Cmd(CmdUploadImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) DeleteUser(param UserId) (err error) {
	var payload []byte
	payload = append(payload, param.UserIdHeb, param.UserIdLeb)
	packet := ps.Builder().Cmd(CmdDeleteUser).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Capture() (err error) {
	packet := ps.Builder().Cmd(CmdCapture).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) SnapImage(param MessageSnapImageData) (err error) {
	var payload []byte
	payload = append(payload, param.ImageCounts, param.StartNumber)
	packet := ps.Builder().Cmd(CmdSnapImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetSavedImage(imageNumber byte) (err error) { // 待传输图片的编号
	var payload []byte
	payload = append(payload, imageNumber)
	packet := ps.Builder().Cmd(CmdGetSavedImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Reset() (err error) {
	packet := ps.Builder().Cmd(CmdReset).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetStatus() (err error) {
	packet := ps.Builder().Cmd(CmdGetStatus).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Verify(param MessageVerifyData) (err error) {
	var payload []byte
	payload = append(payload, param.PowerDownRightAway, param.Timeout)
	packet := ps.Builder().Cmd(CmdVerify).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) DelAll() (err error) {
	packet := ps.Builder().Cmd(CmdDeleteAll).Build()
	err = ps.send(packet)
	return
}
