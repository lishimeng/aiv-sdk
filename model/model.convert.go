package model

func (rc ResultCode) Convert() (rr ReplyResult) {
	switch rc {
	case MrcSuccess:
		rr = MrSuccess
	case MrcRejected:
		rr = MrRejected
	case MrcAborted:
		rr = MrAborted
	case MrcFailed4Camera:
		rr = MrFailed4Camera
	case MrcFailed4UnknownReason:
		rr = MrFailed4UnknownReason
	case MrcFailed4InvalidParam:
		rr = MrFailed4InvalidParam
	case MrcFailed4NoMemory:
		rr = MrFailed4NoMemory
	case MrcFailed4UnknownUser:
		rr = MrFailed4UnknownUser
	case MrcFailed4MaxUser:
		rr = MrFailed4MaxUser
	case MrcFailed4FaceEnrolled:
		rr = MrFailed4FaceEnrolled
	case MrcFailed4LiveNessCheck:
		rr = MrFailed4LiveNessCheck
	case MrcFailed4Timeout:
		rr = MrFailed4Timeout
	case MrcFailed4Authorization:
		rr = MrFailed4Authorization
	case MrcFailed4CameraFov:
		rr = MrFailed4CameraFov
	case MrcFailed4CameraQua:
		rr = MrFailed4CameraQua
	case MrcFailed4CameraStru:
		rr = MrFailed4CameraStru
	case MrcFailed4BootTimeout:
		rr = MrFailed4BootTimeout
	case MrcFailed4ReadFile:
		rr = MrFailed4ReadFile
	case MrcFailed4WriteFile:
		rr = MrFailed4WriteFile
	case MrcFailed4NoEncrypt:
		rr = MrFailed4NoEncrypt
	}
	return
}
