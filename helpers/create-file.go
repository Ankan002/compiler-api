package helpers

import (
	"github.com/gofiber/fiber/v2/utils"
	"io/ioutil"
	"log"
	"os"
)

type FileCreationResponse struct {
	Success  bool
	FileName string
	Error    string
}

func MakeDirectory() bool {
	err := os.Mkdir("code", 0755)

	if err != nil {
		log.Println("Directory creation failed....")

		return false
	}

	return true
}

func CreateFile(code string, language string) FileCreationResponse {
	if _, directorySearchError := os.Stat("code"); os.IsNotExist(directorySearchError) {
		isDirectoryCreated := MakeDirectory()

		if !isDirectoryCreated {
			return FileCreationResponse{
				Success: false,
				Error:   "Directory Creation Failed...",
			}
		}
	}

	fileName := utils.UUIDv4() + "." + language

	writeError := ioutil.WriteFile("code/"+fileName, []byte(code), 0664)

	if writeError != nil {
		return FileCreationResponse{
			Success: false,
			Error:   "File could not be created...",
		}
	}

	return FileCreationResponse{
		Success:  true,
		FileName: fileName,
	}
}
