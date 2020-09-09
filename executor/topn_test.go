package executor

import (
	"testing"
)

func TestT(t *testing.T) {
}
func TestNew(t *testing.T) {
	datas := GetRemoteData(1000, 1100)
	topNExec := NewTopNExec(datas)

	topNExec.Get(10)
}
