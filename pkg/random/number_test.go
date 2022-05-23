package random

import (
	"golayout/pkg/time"
	"testing"
)

func TestRandomNumber(t *testing.T) {
	if Int() == Int() {
		t.Error("Int() should be different")
	}
	n := time.NowUnixNano()
	if Intn(int(n)) == Intn(int(n)) {
		t.Error("Intn() should be different")
	}
}
