package aiv_sdk

func (ps *PacketSdk) IdConfigBaudRate(param BaudRate) (err error) {
	var payload []byte
	payload = append(payload, byte(param))
	packet := ps.Builder().Cmd(CmdIdConfigBaudRate).Body(payload).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) IdSetReleaseEncKey(param []byte) (err error) {
	var payload []byte
	payload = append(payload, param...) // 加密序列 16byte
	packet := ps.Builder().Cmd(CmdIdSetReleaseEncKey).Body(payload).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) IdSetDebugEncKey(param []byte) (err error) {
	var payload []byte
	payload = append(payload, param...) // 加密序列 16byte
	packet := ps.Builder().Cmd(CmdIdSetDebugEncKey).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdOtaHeader(param MessageOtaHeaderData) (err error) {
	var payload []byte
	payload = append(payload, param.TotalSize...)
	payload = append(payload, param.PacketNumber...)
	payload = append(payload, param.PacketSize...)
	payload = append(payload, param.Md5...)
	packet := ps.Builder().Cmd(CmdIdOtaHeader).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdOtaPacket(param MessageOtaPacketData) (err error) {
	var payload []byte
	payload = append(payload, param.PacketId...)
	payload = append(payload, param.PacketSize...)
	payload = append(payload, param.Data...)
	packet := ps.Builder().Cmd(CmdIdOtaPacket).Body(payload).Build()
	err = ps.send(packet)
	return
}
func (ps *PacketSdk) IdStartOta(primary, secondary, revision byte) (err error) { // 固件版本号
	var payload []byte
	payload = append(payload, primary, secondary, revision)
	packet := ps.Builder().Cmd(CmdIdStartOta).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdStopOta() (err error) {
	packet := ps.Builder().Cmd(CmdIdStopOta).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdGetOtaStatus() (err error) {
	packet := ps.Builder().Cmd(CmdIdGetOtaStatus).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdGetAllUserId() (err error) {
	packet := ps.Builder().Cmd(CmdIdGetAllUserid).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) IdEnrollItg(param MessageEnRollItg) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.EnrollType, param.EnableDuplicate, param.Timeout)
	payload = append(payload, 0x00, 0x00, 0x00) // resv
	packet := ps.Builder().Cmd(CmdIdEnrollItg).Body(payload).Build()
	err = ps.send(packet)
	return
}
