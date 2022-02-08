package device

import "github.com/lishimeng/aiv-sdk/model"

func (ps *PacketSdk) IdConfigBaudRate(param model.BaudRate) (err error) {
	var payload []byte
	payload = append(payload, byte(param))
	packet := ps.Builder().Cmd(model.CmdIdConfigBaudRate).Body(payload).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) IdSetReleaseEncKey(param []byte) (err error) {
	var payload []byte
	payload = append(payload, param...) // 加密序列 16byte
	packet := ps.Builder().Cmd(model.CmdIdSetReleaseEncKey).Body(payload).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) IdSetDebugEncKey(param []byte) (err error) {
	var payload []byte
	payload = append(payload, param...) // 加密序列 16byte
	packet := ps.Builder().Cmd(model.CmdIdSetDebugEncKey).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdOtaHeader(param model.MessageOtaHeaderData) (err error) {
	var payload []byte
	payload = append(payload, param.TotalSize...)
	payload = append(payload, param.PacketNumber...)
	payload = append(payload, param.PacketSize...)
	payload = append(payload, param.Md5...)
	packet := ps.Builder().Cmd(model.CmdIdOtaHeader).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdOtaPacket(param model.MessageOtaPacketData) (err error) {
	var payload []byte
	payload = append(payload, param.PacketId...)
	payload = append(payload, param.PacketSize...)
	payload = append(payload, param.Data...)
	packet := ps.Builder().Cmd(model.CmdIdOtaPacket).Body(payload).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) IdStartOta(primary, secondary, revision byte) (err error) { // 固件版本号
	var payload []byte
	payload = append(payload, primary, secondary, revision)
	packet := ps.Builder().Cmd(model.CmdIdStartOta).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdStopOta() (err error) {
	packet := ps.Builder().Cmd(model.CmdIdStopOta).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdGetOtaStatus() (err error) {
	packet := ps.Builder().Cmd(model.CmdIdGetOtaStatus).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdGetAllUserId() (err error) {
	packet := ps.Builder().Cmd(model.CmdIdGetAllUserid).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdEnrollItg(param model.MessageEnRollItg) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.EnrollType, param.EnableDuplicate, param.Timeout)
	payload = append(payload, 0x00, 0x00, 0x00) // resv
	packet := ps.Builder().Cmd(model.CmdIdEnrollItg).Body(payload).Build()
	err = ps.send(packet)
	return
}
