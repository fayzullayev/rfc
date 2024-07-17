package utils

import (
	"log"
	"os"
)

func LogErrorAndCleanup(err error, folderPath string) {
	log.Println("Error:", err)
	if removeErr := os.RemoveAll(folderPath); removeErr != nil {
		log.Fatal("Unable to remove component folder: ", folderPath)
	}
}
