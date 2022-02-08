package device

import "github.com/lishimeng/aiv-sdk/model"

// OnReady
// 模块已准备好
type OnReady func()

// OnUnknownErr
// 未知错误
type OnUnknownErr func()

// OnOtaDone
// OTA 升级完毕
type OnOtaDone func(success bool)

// OnEyeStateCallback
// 解锁过程中睁闭眼状态
type OnEyeStateCallback func()

// OnFaceStateCallback
// 算法执行成功，并且返回人脸信息
type OnFaceStateCallback func(state model.FaceStateWrapper)
