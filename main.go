package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type ComponentTemplData struct {
	Name string
}

func main() {
	currentFolder, err := os.Getwd()
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}

	args := os.Args
	if len(args) != 2 {
		log.Fatal("Folder name is required: ")
	}

	folder := args[1]
	if len(folder) < 4 {
		log.Fatal("Folder name is too short")
	}

	fullPath := path.Join(currentFolder, folder)
	err = os.Mkdir(fullPath, os.ModePerm)

	if err != nil {
		log.Fatal("Folder creation failed: ", err)
	}

	names := strings.Split(folder, "-")

	for i, name := range names {
		newName := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])

		names[i] = newName
	}

	componentName := strings.Join(names, "")

	fmt.Printf("Component name: %s\n", componentName)

	file, err := os.Create(path.Join(fullPath, folder+".tsx"))

	if err != nil {
		log.Fatal("File creation failed: ", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("File creation failed: ", err)
		}
	}(file)

	componentTemplData := ComponentTemplData{
		Name: componentName,
	}

	var tmplFile = "./templates/component.templ"
	tmpl, err := template.New("").ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(file, componentTemplData)
	if err != nil {
		panic(err)
	}

}
