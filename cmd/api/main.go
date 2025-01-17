package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		product := r.URL.Query().Get("product")
		id := r.URL.Query().Get("id")
		if product != "" {
			w.Write([]byte(product + " " + id))
		} else {
			w.Write([]byte(""))
		}

	})

	r.Get("/{productName}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "productName")
		w.Write([]byte(param))
	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		obj := map[string]string{"message": "success"}
		b, _ := json.Marshal(obj)
		w.Write(b)
	})

	http.ListenAndServe(":3000", r)
}
