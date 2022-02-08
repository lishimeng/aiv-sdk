package model

const (
	CmdReady              = 0x00 // 这个信号不会发送, 首次接收时判断用
	CmdReset              = 0x10
	CmdGetStatus          = 0x11
	CmdVerify             = 0x12
	CmdEnroll             = 0x13
	CmdCapture            = 0x14
	CmdSnapImage          = 0x16
	CmdGetSavedImage      = 0x17
	CmdUploadImage        = 0x18
	CmdDeleteUser         = 0x20
	CmdDeleteAll          = 0x21
	CmdGetUserInfo        = 0x22
	CmdFaceReset          = 0x23
	CmdIdGetAllUserid     = 0x24
	CmdIdEnrollItg        = 0x26
	CmdGetVersion         = 0x30
	CmdIdStartOta         = 0x40
	CmdIdStopOta          = 0x41
	CmdIdGetOtaStatus     = 0x42
	CmdIdOtaHeader        = 0x43
	CmdIdOtaPacket        = 0x44
	CmdInitEncryption     = 0x50
	CmdIdConfigBaudRate   = 0x51
	CmdIdSetReleaseEncKey = 0x52
	CmdIdSetDebugEncKey   = 0x53
	CmdGetLogFile         = 0x60
	CmdUploadLogFile      = 0x61
	CmdPowerDown          = 0xED
	CmdDemoMode           = 0xFE
	CmdTransFilePacket    = 0x90
	CmdEnrollFromImage    = 0x91
	CmdGetFeatureInfo     = 0x92
	CmdUploadFeature      = 0x93
	CmdEnrollWithFeature  = 0x94
	CmdEnrollSingle       = 0x1D
)

const (
	FaceDirectionUp        FaceDir = 1 << 4
	FaceDirectionDown      FaceDir = 1 << 3
	FaceDirectionLeft      FaceDir = 1 << 2
	FaceDirectionRight     FaceDir = 1 << 1
	FaceDirectionMiddle    FaceDir = 1 << 0
	FaceDirectionUndefined FaceDir = 0x00
)

const (
	MR_SUCCESS               = 0  // success
	MR_REJECTED              = 1  // module rejected this command
	MR_ABORTED               = 2  // algo aborted
	MR_FAILED4_CAMERA        = 4  // camera open failed
	MR_FAILED4_UNKNOWNREASON = 5  // UNKNOWN_ERROR
	MR_FAILED4_INVALIDPARAM  = 6  // invalid param
	MR_FAILED4_NOMEMORY      = 7  // no enough memory
	MR_FAILED4_UNKNOWNUSER   = 8  // Unknown user
	MR_FAILED4_MAXUSER       = 9  // exceed maximum user number
	MR_FAILED4_FACEENROLLED  = 10 // this face has been enrolled
	MR_FAILED4_LIVENESSCHECK = 12 // liveness check failed
	MR_FAILED4_TIMEOUT       = 13 // exceed the time limit
	MR_FAILED4_AUTHORIZATION = 14 // authorization failed
	MR_FAILED4_CAMERAFOV     = 15 // camera fov test failed
	MR_FAILED4_CAMERAQUA     = 16 // camera quality test failed
	MR_FAILED4_CAMERASTRU    = 17 // camera structure test failed
	MR_FAILED4_BOOT_TIMEOUT  = 18 // boot up timeout
	MR_FAILED4_READ_FILE     = 19 // read file failed
	MR_FAILED4_WRITE_FILE    = 20 // write file failed
	MR_FAILED4_NO_ENCRYPT    = 21 // encrypt must be set
)

const (
	UNLOCK_NORMAL          UnlockStatus = 200
	UNLOCK_WITH_EYES_CLOSE UnlockStatus = 204
)

const (
	MS_STANDBY MessageStatus = 0
	MS_BUSY    MessageStatus = 1
	MS_ERROR   MessageStatus = 2
	MS_INVALID MessageStatus = 3
)

const (
	MsgReply MessageType = 0x00
	MsgNote  MessageType = 0x01
	MsgImage MessageType = 0x02
)

const (
	NID_READY        = 0 // 模块已准备好
	NID_FACE_STATE   = 1 // 算法执行成功，并且返回人脸信息 s_note_data_face
	NID_UNKNOWNERROR = 2 // 未知错误
	NID_OTA_DONE     = 3 // OTA 升级完毕
	NID_EYE_STATE    = 4 // 解锁过程中睁闭眼状态
)

const (
	FACE_STATE_NORMAL                         FaceState = 0  // 人脸正常
	FACE_STATE_NOFACE                         FaceState = 1  // 未检测到人脸
	FACE_STATE_TOOUP                          FaceState = 2  // 人脸太靠近图片上边沿，未能录入
	FACE_STATE_TOODOWN                        FaceState = 3  // 人脸太靠近图片下边沿，未能录入
	FACE_STATE_TOOLEFT                        FaceState = 4  // 人脸太靠近图片左边沿，未能录入
	FACE_STATE_TOORIGHT                       FaceState = 5  // 人脸太靠近图片右边沿，未能录入
	FACE_STATE_FAR                            FaceState = 6  // 人脸距离太远，未能录入
	FACE_STATE_CLOSE                          FaceState = 7  // 人脸距离太近，未能录入
	FACE_STATE_EYEBROW_OCCLUSION              FaceState = 8  // 眉毛遮挡
	FACE_STATE_EYE_OCCLUSION                  FaceState = 9  // 眼睛遮挡
	FACE_STATE_FACE_OCCLUSION                 FaceState = 10 // 脸部遮挡
	FACE_STATE_DIRECTION_ERROR                FaceState = 11 // 录入人脸方向错误
	FACE_STATE_EYE_CLOSE_STATUS_OPEN_EYE      FaceState = 12 // 在闭眼模式检测到睁眼状态
	FACE_STATE_EYE_CLOSE_STATUS               FaceState = 13 // 闭眼状态
	FACE_STATE_EYE_CLOSE_STATUS_UNKOWN_STATUS FaceState = 14 // 在闭眼模式检测中无法判定睁眼状态
)

const (
	BR115200  BaudRate = 0x01
	BR230400  BaudRate = 0x02
	BR460800  BaudRate = 0x03
	BR1500000 BaudRate = 0x04
)

const (
	DemoModeEnable  DemoMode = 0x01
	DemoModeDisable DemoMode = 0x00
)

const (
	WorkStateEnroll = 1
	WorkStateVerify = 2
	WorkStateIdle   = 0
)

var (
	DevStateIdle    = DeviceState{State: 0, Desc: "IDLE"}
	DevStateBusy    = DeviceState{State: 1, Desc: "BUSY"}
	DevStateError   = DeviceState{State: 2, Desc: "ERROR"}
	DevStateInvalid = DeviceState{State: 3, Desc: "INVALID"}
)
