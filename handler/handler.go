package handler

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func Start(port int) (err error) {
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":"+string(port), nil)
	return
}
