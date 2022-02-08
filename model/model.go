package model

// MessageType
// M->H 上报消息类型
type MessageType byte

type MessageStatus byte

type UnlockStatus int

type FaceState int

type DeviceState struct {
	State int
	Desc  string
}

// BaudRate
// 波特率index
type BaudRate byte

type DemoMode byte

// FaceStateWrapper
// left:人脸框距离图片最左侧的距离（负数表示人脸框已超出图片最左侧）
// top:人脸框距离图片最上方的距离（负数表示人脸框已超出图片上方）
// right:人脸框距离图片最右侧的距离（负数表示人脸框已超图片最右侧）
// bottom:人脸框距离图片最下方的距离（负数表示人脸框已超出图片最下方）
// yaw, pitch, roll 表示的旋转方向分别如下图所示。
// 其中
//     yaw 为负表示左转头，yaw 为正表示 右转头；
//     pitch 为负表示向上抬头，pitch 为正表示向下低头；
//     roll 为负表示向右歪头，roll 为正 表示向左歪头
type FaceStateWrapper struct {
	FaceState
	Left   uint16
	Top    uint16
	Right  uint16
	Bottom uint16
	Yaw    uint16
	Pitch  uint16
	Roll   uint16
}

type ReplyBody struct {
	Mid    byte
	Result byte
	Data   []byte
}

type MessageNote struct {
	Nid  byte
	Data []byte
}

type UserId struct {
	UserIdHeb byte
	UserIdLeb byte
}

type UserInfo struct {
	UserId
	UserName []byte
	Admin    byte
}

type MessageReplyVerifyData struct {
	UserInfo
	UnlockStatus
}

type MessageVerifyData struct {
	PowerDownRightAway byte // 解锁成功后是否立刻断 电(yes:1 no:0)
	Timeout            byte
}

type MessageCode int

type FaceDir byte

type MessageEnRoll struct {
	Admin    byte   // 是否设置为管理员 (yes:1 no:0 )
	UserName []byte // 录入用户姓名
	Timeout  byte
	FaceDir  // 用户需要录入的方向， 具体参数如表 4-7 所 示
}

type MessageEnRollItg struct {
	Admin           byte   // 是否设置为管理员 (yes:1 no:0 )
	UserName        []byte // 录入用户姓名
	Timeout         byte
	FaceDir                // 用户需要录入的方向， 具体参数如表 4-7 所 示
	EnrollType      byte   // 录入方式（交互式录 入:0 单张录入:1）
	EnableDuplicate byte   // 是否允许用户重复录入 （不允许重复录入:0 允许重复录入:1）
	Resv            []byte // 暂不支持 3byte
}

type MessageReplyEnRoll struct {
	UserId
	FaceDir byte
}

type MessageSnapImageData struct {
	ImageCounts byte // 抓拍图片的数量
	StartNumber byte // 抓拍图片的起始 编号(1-30)
}

type MessageUploadImageData struct {
	UploadImageOffset []byte // 待上传图片的偏移量
	UploadImageSize   []byte // 此次上传图片的大小， 最大 4K
}

type MessageOtaHeaderData struct {
	TotalSize    []byte // 固件大小，高位在前 b0 b1 b2 b3
	PacketNumber []byte // 总包数，高位在前 b0 b1 b2 b3
	PacketSize   []byte // 每包大小，高位在前 b0 b1
	Md5          []byte // 固件 MD5 值
}

type MessageOtaPacketData struct {
	PacketId   []byte // 包序号 b0 b1
	PacketSize []byte // 包大小，高位在前 b0 b1
	Data       []byte // 分包内容
}

type MessageInitEncryptionData struct {
	Seed       []byte // 主控发送一个随机数 b0 b1 b2 b3
	Mode       byte   // 加密模式
	CreateTime []byte // 当前时区的秒数(1970 年至今) b0 b1 b2 b3
}
type MessageUploadLogFileData struct {
	LogFileOffset []byte // 待上传日志的偏移量 b0 b1 b2 b3
	LogFileSize   []byte // 此次上传日志的大小， 最大 4K b0 b1 b2 b3
}

type MessageTransFileData struct {
	StoreType  byte   // 0,目前仅支持存放在内 存
	FileSize   []byte // 文件大小 b0 b1 b2 b3
	Offset     []byte // 本包数据偏移值 b0 b1 b2 b3
	PacketSize []byte //  本包数据大小 b0 b1
	Data       []byte // 数据
}

type MessageUploadFaceFeature struct {
	Size   []byte // 上传数据的大小 b0 b1 b2 b3
	Offset []byte // 上传数据的偏移值 b0 b1 b2 b3
}
