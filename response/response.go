package response

import (
	"net/http"
	"encoding/json"
	"fmt"
	"git.boois.cn/d01/git_repo/boois_utils.git/result"
)

func ResponseRes(w http.ResponseWriter, res result.ReturnResult) {
	bytes, _ := json.Marshal(res)

	w.Write(bytes)
}
func Response(w http.ResponseWriter, code int, msg string, info string, debug string, result interface{}) {
	bytes, _ := json.Marshal(result.ReturnResult{
		Code:code,
		Msg:msg,
		Info:info,
		Debug:debug,
		Result:result,
	})

	w.Write(bytes)

}