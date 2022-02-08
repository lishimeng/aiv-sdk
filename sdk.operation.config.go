package aiv_sdk

import (
	"fmt"
	"github.com/lishimeng/aiv-sdk/model"
	"github.com/lishimeng/aiv-sdk/util"
	"time"
)

func (s *SerialSdk) SetDefaultCallback(p interface{}) {
	fmt.Println(p)
}

func (s *SerialSdk) Reset() {
	err := s.dev.Reset()
	if err != nil {
		fmt.Println(err) // TODO?
	}
}

func (s *SerialSdk) GetStatus(callback func(state model.DeviceState)) {
	if callback != nil {
		// TODO set callback
	}
	s.dev.SetReplyHandler(func(midMatch, success bool, body model.ReplyBody) {

	})
	err := s.dev.GetStatus()
	if err != nil {
		fmt.Println(err) // TODO?
	}
}

func (s *SerialSdk) GetVersion(callback Callback) {

	err := s.dev.GetVersion()
	if err != nil {
		fmt.Println(err) // TODO?
	}
}

func (s *SerialSdk) InitEncryption(password string, callback Callback) {

	var timer *time.Timer

	// seed
	// mode
	// timestamp
	s.dev.SetReplyHandler(func(midMatch, success bool, body model.ReplyBody) {
		if timer != nil {
			timer.Stop() // 停止超时
		}
		s.dev.SetReplyHandler(func(midMatch, success bool, reply model.ReplyBody) {
			callback(success, nil) // TODO
		})
		// TODO if success
		e := s.dev.InitEncryption(model.MessageInitEncryptionData{
			Seed:       []byte{0x01, 0x02, 0x03, 0x04},
			Mode:       0x01, // TODO ? 文档中没提 能枚举哪些值
			CreateTime: util.TimeNow(),
		})
		if e != nil {
			fmt.Println(e) // TODO?
		}
	})
	err := s.dev.IdSetReleaseEncKey([]byte(password)) // 设置加密序列(16byte)

	if err != nil {
		fmt.Println(err) // TODO?
		return
	}
	timer = time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
		// TODO 执行超时了
	}
}

func (s *SerialSdk) SetBaudRate(rate model.BaudRate) {

	s.dev.SetReplyHandler(nil)
	err := s.dev.IdConfigBaudRate(rate)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *SerialSdk) PowerDown(callback Callback) {
	err := s.dev.PowerDown()
	if err != nil {
		fmt.Println(err)
	}
}
