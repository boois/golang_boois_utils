package boois_log

import (
	"testing"
)

func TestAll(t *testing.T) {
	SetLv(LV_WARN)

	I("info")
	W("warn")
	D("debug")
	E("error")
}
