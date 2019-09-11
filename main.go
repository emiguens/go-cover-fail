package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emiguens/go-cover-fail/pkg1"
)

func greet(w http.ResponseWriter, r *http.Request) {
	_ = pkg1.HasHeader(r, "x-ttl")
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
