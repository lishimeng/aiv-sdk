package communication

import (
	"bytes"
	"testing"
)

func Test001(t *testing.T) {

	var b = bytes.Buffer{}

	t.Log("write string")
	b.WriteString("this_is_a_test")

	t.Log(b.String())
	var c, err = b.ReadByte()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(c))
	b.WriteByte(c)
	c, err = b.ReadByte()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(c))
	t.Log(b.String())
	tmp := b.Next(2)
	t.Log("print 2")
	t.Log(string(tmp))
	t.Log(b.String())
	t.Log(string(b.Bytes()[:2]))
	for {
		if b.Len() > 0 {
			c, err = b.ReadByte()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(c))
		} else {
			break
		}
	}
	t.Log(b.String())
}

func Test002(t *testing.T) {
	var a []int
	var b []int
	a = append(a, 1, 2, 3)
	b = append(b, a[:3]...)
	t.Log("print")
	for _, v := range a {
		t.Log(v)
	}

	a = a[0:0]
	t.Log("print")
	for _, v := range b {
		t.Log(v)
	}
}
