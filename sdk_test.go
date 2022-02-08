package aiv_sdk

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	var ctx = context.TODO()
	var s Sdk = New(ctx)
	t.Logf("%v", s)
}
