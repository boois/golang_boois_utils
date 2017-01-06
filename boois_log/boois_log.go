package boois_log

import (
	"log"
	"os"
	"fmt"
)

const (
	LV_FATAL = 5
	LV_ERROR = 4
	LV_WARN = 3
	LV_INFO = 2
	LV_DEBUG = 1
)

var (
	ins_F *log.Logger
	ins_E *log.Logger
	ins_W *log.Logger
	ins_I *log.Logger
	ins_D *log.Logger
)
//默认使用warn级别
var lv = LV_WARN

func SetLv(lv1 int) {
	lv = lv1
}
func init() {
	ins_F = log.New(os.Stderr, "[FATAL] ", log.LstdFlags | log.Llongfile)
	ins_E = log.New(os.Stderr, "[ERROR] ", log.LstdFlags | log.Llongfile)

	ins_W = log.New(os.Stdout, "[WARN] ", log.LstdFlags )
	ins_I = log.New(os.Stdout, "[INFO] ", log.LstdFlags  | log.Llongfile)
	ins_D = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags | log.Llongfile)
}
func F(v ...interface{}) {
	if LV_FATAL >= lv {
		ins_F.Output(2, fmt.Sprintln(v...))
	}
}
func E(v ...interface{}) {
	if LV_ERROR >= lv {
		ins_E.Output(2, fmt.Sprintln(v...))
	}
}
func W(v ...interface{}) {
	if LV_WARN >= lv {
		ins_W.Output(2, fmt.Sprintln(v...))
	}
}
func I(v ...interface{}) {
	if LV_INFO >= lv {
		ins_I.Output(2, fmt.Sprintln(v...))
	}
}
func D(v ...interface{}) {
	if LV_DEBUG >= lv {
		ins_D.Output(2, fmt.Sprintln(v...))
	}
}

