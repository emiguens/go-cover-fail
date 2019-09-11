package main

import (
	"net/http"

	"github.com/emiguens/go-cover-fail/pkg1"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_ = pkg1.HasHeader(r, "x-ttl")
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}
