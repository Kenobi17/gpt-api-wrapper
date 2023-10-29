package main

import (
	"fmt"
	"net/http"

	"github.com/jba/muxpatterns"
)

func main() {
	mux := muxpatterns.NewServeMux()

	mux.HandleFunc("GET /path", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "got path\n")
	})

	http.ListenAndServe("localhost:8080", mux)
}
