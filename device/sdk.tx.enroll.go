package device

import "github.com/lishimeng/aiv-sdk/model"

func (ps *PacketSdk) EnrollWithFeature(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnrollWithFeature).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) EnrollSingle(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnrollSingle).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) EnrollFromImage(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnrollFromImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

func (ps *PacketSdk) Enroll(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnroll).Body(payload).Build()
	err = ps.send(packet)
	return
}
