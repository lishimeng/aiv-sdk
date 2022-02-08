package model

type UserInfoH struct {
	UserId   string
	UserName string
	Admin    bool
}

func (h *UserInfoH) FromDevice(m UserInfo) *UserInfoH {
	h.UserName = string(m.UserName)
	// userid -> hex
	return h
}

type EnrollMode int

const (
	EnrollH           EnrollMode = 0x13 // 交互式录入用户
	EnrollWithFeature EnrollMode = 0x94 // 通过特征注册用户
	EnrollSingle      EnrollMode = 0x1D // 通过单张人脸注册用户
)

type Version struct {
	v byte

	str string
}

func (v Version) String() string {
	return v.str
}

func (v Version) V() int {
	return int(v.v)
}
