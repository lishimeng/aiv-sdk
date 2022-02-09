package aiv_sdk

import (
	"fmt"
	"github.com/lishimeng/aiv-sdk/model"
	"github.com/lishimeng/aiv-sdk/util"
	"time"
)

func (s *SerialSdk) InitEncryption(password string, callback Callback) {

	// set password(encrypt serial)
	// set seed / timestamp
	// calc session: encrypt key (use password+seed)

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
