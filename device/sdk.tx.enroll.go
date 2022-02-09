package device

import "github.com/lishimeng/aiv-sdk/model"

// EnrollWithFeature
// ID_ENROLL_WITH_FEATURE:注册
//   特征注册时,模组固件会解析特征文件,并以特征文件中包含的信息进行注册.
//   为使各个注册方法的协议保持一致，所以特征注册协议的参数中也包含了 user_name, face_direction等信息.
//   但因为特征文件中已经包含了用户的所有信息,所以这些参数在模组的固件中没有进行任何处理,其作用仅为填充.
func (ps *PacketSdk) EnrollWithFeature(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnrollWithFeature).Body(payload).Build()
	err = ps.send(packet)
	return
}

// EnrollSingle
// 单张图片注册
func (ps *PacketSdk) EnrollSingle(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnrollSingle).Body(payload).Build()
	err = ps.send(packet)
	return
}

// EnrollFromImage
// 下发单张图片注册
func (ps *PacketSdk) EnrollFromImage(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnrollFromImage).Body(payload).Build()
	err = ps.send(packet)
	return
}

// Enroll
// 交互式注册
func (ps *PacketSdk) Enroll(param model.MessageEnRoll) (err error) {
	var payload []byte
	payload = append(payload, param.Admin)
	payload = append(payload, param.UserName...)
	payload = append(payload, byte(param.FaceDir), param.Timeout)
	packet := ps.Builder().Cmd(model.CmdEnroll).Body(payload).Build()
	err = ps.send(packet)
	return
}
