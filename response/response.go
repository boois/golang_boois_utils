package response

import (
	"net/http"
	"encoding/json"
	"fmt"
	"result"
)

func ResponseRes(w http.ResponseWriter, res result.ReturnResult) {
	bytes, _ := json.Marshal(res)

	fmt.Fprint(w, string(bytes))
}
func Response(w http.ResponseWriter, code int, msg string, info string, debug string, result interface{}) {
	bytes, _ := json.Marshal(result.ReturnResult{
		Code:code,
		Msg:msg,
		Info:info,
		Debug:debug,
		Result:result,
	})

	fmt.Fprint(w, string(bytes))
}