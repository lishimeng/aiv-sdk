package device

import "github.com/lishimeng/aiv-sdk/model"

func (ps *PacketSdk) StartVerify(powerOff bool) (err error) {

	ps.workState = model.WorkStateVerify
	ps.handleReply = ps.verifyProcess

	var po byte = 0x00
	if powerOff {
		po = 0x01
	}
	err = ps.Verify(model.MessageVerifyData{
		PowerDownRightAway: po,
		Timeout:            10,
	})
	// if err cancel
	if err != nil {
		ps.StopVerify()
	}
	return
}

func (ps *PacketSdk) StopVerify() {

	ps.workState = model.WorkStateIdle
	ps.handleReply = nil
}

func (ps *PacketSdk) verifyProcess(midMatch, success bool, body model.ReplyBody) {

	if body.Result == 0x01 {
		// TODO
	}
	// TODO success?
	ps.StopVerify()
}
