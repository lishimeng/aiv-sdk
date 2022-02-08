package aiv_sdk

// ① 主控发送 ID_START_OTA（加密通信）通知 模块 进入 OTA 模式；
// ② 模块 反馈 OK（不加密通信，后面的都是不加密模式），表示进入 了 OTA 模式；
// ③ 主控发送 ID_CONFIG_BAUDRATE 配置更高波特率， 模块 反馈 OK 后，主机同步设置相同波特率。延时 100ms 后才能开始串口数据传输。
//（该配置过程可选，波特率默认是 115200
// ④ 主控发送 ID_GET_OTA_STATUS
// ⑤ 模块反馈 OTA 状态和起始包序号
// ⑥ 主控发送 ID_OTA_HEADER，把固件相关信息发送到模块
// ⑦ 模块反馈 OK
// ⑧ 主控循环发送升级包 ID_OTA_PACKET，收到 OK 后，继续下一包。
// ⑨ 模块 收完固件后，开始升级。
// ⑩ 模块 升级结束反馈 NID_OTA_DONE。
//具体流程参照：
//1) “智能人脸模组 OTA 升级方案 V1.0.pptx”
//2) “智能人脸模组 OTA 方案软件工作流程 .pdf”

// StartOta 开始OTA
func (s *SerialSdk) StartOta() { // 进入OTA

}

func (s *SerialSdk) StopOta() { // 退出OTA

}

func (s *SerialSdk) otaProcess() { // OTA过程

}
