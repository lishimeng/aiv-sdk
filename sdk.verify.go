package aiv_sdk

func (ps *PacketSdk) StartVerify(powerOff bool) (err error) {

	ps.workState = WorkStateVerify
	ps.handleReply = ps.verifyProcess

	var po byte = 0x00
	if powerOff {
		po = 0x01
	}
	err = ps.Verify(MessageVerifyData{
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

	ps.workState = WorkStateIdle
	ps.handleReply = nil
}

func (ps *PacketSdk) verifyProcess(body ReplyBody) {

	// TODO success?
	ps.StopVerify()
}
