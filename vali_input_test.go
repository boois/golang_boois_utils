package boois_utils

import (
	"testing"
	"net/http"
	"log"
	"fmt"
)

func test(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form)
	ret:=Validete(r,
		`app_key -n -t guid`,
		`page -n -t +int`,
		`page_size -n -t +int`,
		`price -n -t +float|zero`,
	)
	fmt.Print(ret)
	w.Write([]byte(ret.Info))

}

func TestValidete(t *testing.T) {
	http.HandleFunc("/test", test)
	err := http.ListenAndServe("127.0.0.1:9875", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
