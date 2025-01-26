package files

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"time"

	"dossantoscarlos.com/goPDF/backend/internal/structs"
)

func CreateFile(file multipart.File, header *multipart.FileHeader) (structs.File, error) {
	var err error

	fileStruct := structs.File{}

	path := "uploads/pdfs/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fileStruct, fmt.Errorf("Erro ao criar o diretório")
		}
	}

	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), header.Filename)
	filePath := path + fileName

	// Salva o arquivo no diretório
	outFile, err := os.Create(filePath)
	if err != nil {
		return fileStruct, fmt.Errorf("Erro ao salvar o arquivo")
	}

	defer outFile.Close()

	// Copia o conteúdo do arquivo enviado para o arquivo criado
	_, err = io.Copy(outFile, file)
	if err != nil {
		return fileStruct, fmt.Errorf("Erro ao salvar o arquivo")
	}

	fileStruct.Path = filePath

	// Resposta de sucesso
	return fileStruct, nil
}
