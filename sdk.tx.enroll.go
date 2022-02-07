package aiv_sdk

func (ps *PacketSdk) EnrollWithFeature(param MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(CmdEnrollWithFeature).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) EnrollSingle(param MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(CmdEnrollSingle).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) EnrollFromImage(param MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(CmdEnrollFromImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Enroll(param MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(CmdEnroll).Body(payload).Build()
	err = ps.send(packet)
	return
}
