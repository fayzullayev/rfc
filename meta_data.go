package main

import (
	"log"
	"os"
	"path"
	"strings"
)

func GetMetaData() FileMetaData {

	currentFolder, err := os.Getwd()
	if err != nil {
		log.Fatal("Something went wrong.: ", err)
	}

	args := os.Args
	if len(args) != 2 {
		log.Fatal("Folder name is required: ")
	}

	componentFolderName := args[1]
	if len(componentFolderName) < 4 {
		log.Fatal("Folder name is too short")
	}

	fullPath := path.Join(currentFolder, componentFolderName)
	err = os.Mkdir(fullPath, os.ModePerm)

	if err != nil {
		log.Fatal("Folder creation failed: ", err)
	}

	names := strings.Split(componentFolderName, "-")

	for i, name := range names {
		newName := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

		names[i] = newName
	}

	componentName := strings.Join(names, "")

	return FileMetaData{
		ComponentName:       componentName,
		ComponentFolderPath: fullPath,
		ComponentFolderName: componentFolderName,
	}
}
