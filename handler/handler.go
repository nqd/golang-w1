package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/nqd/lab1/shortener"
)

func Start(port int, db shortener.Shortener) (err error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// remove first "/"
		alias := trimFirstRune(r.RequestURI)

		redURL, ok := db.Query(alias)

		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Alias '%s' not found", alias)

			return
		}

		http.Redirect(w, r, redURL, http.StatusMovedPermanently)
		return
	})

	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
