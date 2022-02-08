package aiv_sdk

import (
	"context"
	"github.com/lishimeng/aiv-sdk/device"
	"github.com/lishimeng/aiv-sdk/model"
)

type Callback func(bool, interface{})

type Sdk interface {
	SetDefaultCallback(interface{})
	Reset()
	GetStatus(callback func(state model.DeviceState))

	StartOta()
	StopOta()
	otaProcess()
}

type SerialSdk struct {
	ctx context.Context
	dev device.PacketSdk
}

func New(ctx context.Context) Sdk {

	var s Sdk
	var ds = SerialSdk{ctx: ctx}
	s = &ds
	return s
}

func (s *SerialSdk) Connect() {
	// connect serial
}
