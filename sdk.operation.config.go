package aiv_sdk

import (
	"fmt"
	"github.com/lishimeng/aiv-sdk/model"
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
