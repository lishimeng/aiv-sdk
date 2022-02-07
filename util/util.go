package util

func Bytes2Int16(b []byte) int16 {
	return int16(b[0]&0xff) + int16(b[1]&0xff)*256
}

func CalcChk(b []byte, size int16) bool {
	return false
}
