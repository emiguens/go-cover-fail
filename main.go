package main

import (
	"fmt"
	"net/http"

	"github.com/emiguens/go-cover-fail/pkg1"
	"github.com/emiguens/go-cover-fail/pkg2"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		hh1 := pkg1.HasHeader(r, "x-ttl")
		hh2 := pkg2.HasHeader(r, "x-ttl")

		fmt.Fprintf(w, "%t %t", hh1, hh2)

		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
