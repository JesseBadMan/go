package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/form", TestForm)
	http.ListenAndServe(":9090", nil)
}

func TestForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "only support post method")
	}

	r.ParseMultipartForm(100)

	fmt.Fprint(w, "你的Form表单提交了如下信息：")
	fmt.Fprint(w, r.PostForm)
}
