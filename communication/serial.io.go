package communication

import (
	"bytes"
	"context"
	"io"
)

var defaultListener = func(buffer bytes.Buffer) {}

type Conn struct {
	io.ReadWriteCloser
	ctx      context.Context
	buf      bytes.Buffer
	Listener func(buffer bytes.Buffer)
}

func New(ctx context.Context) *Conn {
	c := &Conn{ctx: ctx, Listener: defaultListener}
	return c
}

func (c *Conn) SetRxListener(listener func(buffer bytes.Buffer)) {
	if listener == nil {
		listener = defaultListener
	}
	c.Listener = listener
}

func (c *Conn) Loop() {
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			c.read()
		}
	}
}

func (c *Conn) read() {
	var p []byte
	size, err := c.Read(p)
	if err != nil {
		if err == io.EOF {
			return
		}
		panic(err) // throw err
	}
	if size > 0 { // 读到后再执行操作
		c.buf.Write(p[:size])
		c.Listener(c.buf)
	}
}

func (c *Conn) tx(data []byte) (int, error) {
	return c.Write(data)
}
