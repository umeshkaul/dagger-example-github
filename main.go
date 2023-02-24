package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello from Dagger Go SDK")
	})

	http.ListenAndServe(":9090", nil)
}
