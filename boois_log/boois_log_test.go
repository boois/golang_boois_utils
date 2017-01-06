package boois_log

import (
	"testing"
)

func TestAll(t *testing.T) {
	SetLv(LV_WARN)

	I("info")
	D("debug")
	E("error")
	F("fatal")
	W("warn")
}
