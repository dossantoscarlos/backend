package routers

import (
	"net/http"

	"dossantoscarlos.com/goPDF/backend/cmd/api/handler"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("/upload", handler.Download)
}
