package main

import (
	"github.com/campadrenalin/spiro_identicons/art"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers["Content-Type"] = []string{"image/png"}
	ar := art.NewRequest(r.URL.Path[1:])
	ar.RenderPNG(w)
}
