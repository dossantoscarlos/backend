package routers

import (
	"log"
	"net/http"

	"dossantoscarlos.com/goPDF/backend/cmd/api/handler"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})
	log.Default().Println("Router")

	mux.HandleFunc("/upload", handler.Download)
}
