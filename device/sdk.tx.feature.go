package device

import "github.com/lishimeng/aiv-sdk/model"

// feature(人脸特征值)相关操作, 可以达到备份人脸信息, 转移到其他模块, 免去多个设备重复采集人脸
//在使用特征注册指令前，需要先使用
// ID_GET_FEATURE_INFO: 查询某个用户的特征值文件, 主要拿到文件大小
// ID_UPLOAD_FEATURE: 让模块上传特征值文件(用来备份), 分包上传
// 得到特征文件后，即可使用此特征文件进行注册。
// 用特征注册的第一步需要先下发特征文件，
// ID_TRANS_FILE_PACKET：下发特征值文件到模块
// ID_ENROLL_WITH_FEATURE： 模块得到特征值文件后就可以使用feature注册用户, 过程中使用上一步下发来的特征值文件

func (ps *PacketSdk) GetFeatureInfo(param model.UserId) (err error) {
	// ID_GET_FEATURE_INFO命令可以通过 输入 已知的注册用户的 ID 获取存在 数据 库中此 ID对应的特征信息的大小

	// 得到 特征信息 大小之后，可以 分多次调用 ID_UPLOAD_FEATURE 来获取特征文件
	// ID_UPLOAD_FEATURE 指令中携带的数据 upload_feature_offset[4]表示该次上传 特征信息的 的偏移量
	//（例如偏移量为 0时将从特征文件 开头开始传输），
	// upload_feature_size[4]表示该次上传特征文件 的大小，最大支持 4000 字节。
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
