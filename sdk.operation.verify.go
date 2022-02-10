package aiv_sdk

import (
	"context"
	"fmt"
	"github.com/lishimeng/aiv-sdk/model"
	"time"
)

const (
	defaultVerifyTimeout = 10
)

func (s *SerialSdk) Verify(callback Callback, onFace func(faceData model.FaceStateWrapper), timeout int) {

	var timer *time.Timer
	var ctx, cancel = context.WithCancel(context.Background())
	if timeout < defaultVerifyTimeout {
		timeout = defaultVerifyTimeout
	}
	//
	s.dev.SetNoteHandler(func(messageNote model.MessageNote) {
		if onFace != nil {
			onFace(model.FaceStateWrapper{
				FaceState: 0,
				Left:      0,
				Top:       0,
				Right:     0,
				Bottom:    0,
				Yaw:       0,
				Pitch:     0,
				Roll:      0, // TODO convert
			})
		}
	})
	s.dev.SetReplyHandler(func(midMatch bool, success bool, body model.ReplyBody) {
		if timer != nil {
			timer.Stop()
		}
		cancel()
		if callback != nil {
			callback(success, body) // TODO convert to user struct
		}
	})
	err := s.dev.Verify(model.MessageVerifyData{
		PowerDownRightAway: 0,
		Timeout:            byte(timeout),
	})
	if err != nil {
		fmt.Println(err)
	}

	timer = time.NewTimer(time.Second * (time.Duration(timeout + 1)))
	select {
	case <-ctx.Done():
		// 超时
		return
	case <-timer.C:
		return
	}
}
