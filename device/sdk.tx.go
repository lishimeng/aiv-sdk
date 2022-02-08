package device

import "github.com/lishimeng/aiv-sdk/model"

func (ps *PacketSdk) GetFeatureInfo(param model.UserId) (err error) {
	var payload []byte
	payload = append(payload, param.UserIdHeb, param.UserIdLeb)
	packet := ps.Builder().Cmd(model.CmdGetFeatureInfo).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) UploadFeature(param model.MessageUploadFaceFeature) (err error) {
	var payload []byte
	payload = append(payload, param.Offset...)
	payload = append(payload, param.Size...)
	packet := ps.Builder().Cmd(model.CmdUploadFeature).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetLogFile() (err error) {
	packet := ps.Builder().Cmd(model.CmdGetLogFile).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) UploadLogFile(param model.MessageUploadLogFileData) (err error) {
	var payload []byte
	payload = append(payload, param.LogFileOffset...)
	payload = append(payload, param.LogFileSize...)
	packet := ps.Builder().Cmd(model.CmdUploadLogFile).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) PowerDown() (err error) {
	packet := ps.Builder().Cmd(model.CmdPowerDown).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) DemoMode(enable model.DemoMode) (err error) {
	var payload []byte
	payload = append(payload, byte(enable))
	packet := ps.Builder().Cmd(model.CmdDemoMode).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) TransFilePacket(param model.MessageTransFileData) (err error) {
	var payload []byte
	payload = append(payload, param.StoreType)
	payload = append(payload, param.FileSize...)
	payload = append(payload, param.PacketSize...)
	payload = append(payload, param.Data...)
	packet := ps.Builder().Cmd(model.CmdTransFilePacket).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) InitEncryption(param model.MessageInitEncryptionData) (err error) {
	var payload []byte
	payload = append(payload, param.Seed...)
	payload = append(payload, param.Mode)
	payload = append(payload, param.CreateTime...)
	packet := ps.Builder().Cmd(model.CmdInitEncryption).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetVersion() (err error) {
	packet := ps.Builder().Cmd(model.CmdGetVersion).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) DeleteAll() (err error) {
	packet := ps.Builder().Cmd(model.CmdDeleteAll).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) GetUserInfo(param model.UserId) (err error) {
	var payload []byte
	payload = append(payload, param.UserIdHeb, param.UserIdLeb)
	packet := ps.Builder().Cmd(model.CmdGetUserInfo).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) FaceReset() (err error) {
	packet := ps.Builder().Cmd(model.CmdFaceReset).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) UploadImage(param model.MessageUploadImageData) (err error) {
	var payload []byte
	payload = append(payload, param.UploadImageOffset...)
	payload = append(payload, param.UploadImageSize...)
	packet := ps.Builder().Cmd(model.CmdUploadImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) DeleteUser(param model.UserId) (err error) {
	var payload []byte
	payload = append(payload, param.UserIdHeb, param.UserIdLeb)
	packet := ps.Builder().Cmd(model.CmdDeleteUser).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Capture() (err error) {
	packet := ps.Builder().Cmd(model.CmdCapture).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) SnapImage(param model.MessageSnapImageData) (err error) {
	var payload []byte
	payload = append(payload, param.ImageCounts, param.StartNumber)
	packet := ps.Builder().Cmd(model.CmdSnapImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetSavedImage(imageNumber byte) (err error) { // 待传输图片的编号
	var payload []byte
	payload = append(payload, imageNumber)
	packet := ps.Builder().Cmd(model.CmdGetSavedImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Reset() (err error) {
	packet := ps.Builder().Cmd(model.CmdReset).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) GetStatus() (err error) {
	packet := ps.Builder().Cmd(model.CmdGetStatus).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Verify(param model.MessageVerifyData) (err error) {
	var payload []byte
	payload = append(payload, param.PowerDownRightAway, param.Timeout)
	packet := ps.Builder().Cmd(model.CmdVerify).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) DelAll() (err error) {
	packet := ps.Builder().Cmd(model.CmdDeleteAll).Build()
	err = ps.send(packet)
	return
}
