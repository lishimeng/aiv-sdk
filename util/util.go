package util

func Bytes2Int16(b []byte) int16 {
	return int16(b[0]&0xff) + int16(b[1]&0xff)*256
}

// CalcChk 需传入有效部分
//   有效部分不包括SyncWork和ParityCheck
//   计算方法：各字节按位XOR
func CalcChk(bs []byte, size int16) byte {
	var chk byte
	var i int16
	for i = 0; i < size; i++ {
		b := bs[i]
		if i == 0 {
			chk = b
		} else {
			chk = chk ^ b
		}
	}
	return chk
}

func PacketChk(bs []byte, size int16, parityCheck byte) bool {
	var chk = CalcChk(bs, size)
	return parityCheck == chk
}

// TimeNow 当前时区的秒数(1970 年至今) b0 b1 b2 b3
func TimeNow() (now []byte) {
	// TODO impl
	return
}
