package handler

import (
	"log"
	"net/http"
	"os"
	"strings"

	"dossantoscarlos.com/goPDF/backend/internal/files"
)

func Download(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Default().Println("Método inválido")
		http.Error(w, "Método inválido", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // Limite de 10MB
	if err != nil {
		log.Default().Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Default().Println(err.Error())
		http.Error(w, "Arquivo nao encontrado", http.StatusBadRequest)
		return
	}

	defer file.Close()

	if !strings.HasPrefix(header.Header.Get("Content-Type"), "application/pdf") {
		log.Default().Println("O arquivo enviado não é um PDF")
		http.Error(w, "O arquivo enviado não é um PDF", http.StatusBadRequest)
		return
	}

	fileStruct, err := files.CreateFile(file, header)
	if err != nil {
		log.Default().Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Transfer-Encoding", "binary")

	http.ServeFile(w, r, fileStruct.Path)

	os.Remove(fileStruct.Path)
}
