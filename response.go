package boois_utils

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func ResponseRes(w http.ResponseWriter, res ReturnResult) {
	bytes, _ := json.Marshal(res)

	fmt.Fprint(w, string(bytes))
}
func Response(w http.ResponseWriter, code int, msg string, info string, debug string, result interface{}) {
	bytes, _ := json.Marshal(ReturnResult{
		Code:code,
		Msg:msg,
		Info:info,
		Debug:debug,
		Result:result,
	})

	fmt.Fprint(w, string(bytes))
}