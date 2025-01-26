package main

import (
	"log"
	"net/http"

	"dossantoscarlos.com/goPDF/backend/cmd/api/middleware"
	"dossantoscarlos.com/goPDF/backend/cmd/api/routers"
)

func main() {

	mux := http.NewServeMux()

	middleware := middleware.EnableCORS(mux)

	routers.Router(mux)

	if err := http.ListenAndServe(":8080", middleware); err != nil {
		log.Fatal(err)
	}

}
