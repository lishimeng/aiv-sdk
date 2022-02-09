package model

import "testing"

func TestConvert(t *testing.T) {
	var a = MrSuccess
	var b = ReplyResult{
		ResultCode: 0x00,
		Message:    "成功",
	}

	var match = a.ResultCode == b.ResultCode
	t.Logf("match?: %t", match)
}
