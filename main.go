package main

import (
	"embed"
	"log"
	"os"
)

type FileMetaData struct {
	ComponentFolderName string
	ComponentName       string
	ComponentFolderPath string
}

//go:embed "templates"
var Files embed.FS

func main() {
	metaData := GetMetaData()

	//creating component file
	err := CreateComponentFile(metaData, metaData.ComponentFolderName+".tsx", "component")

	//creating styles file
	err = CreateComponentFile(metaData, "index.ts", "index")

	//creating index file
	err = CreateComponentFile(metaData, "style.ts", "style")

	if err != nil {
		err = os.RemoveAll(metaData.ComponentFolderPath)
		if err != nil {
			log.Fatal("Unable to remove component folder " + metaData.ComponentFolderPath)
		}
	}

}
