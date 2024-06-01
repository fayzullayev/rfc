package main

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type ComponentTemplData struct {
	Name string
	File string
}

//go:embed "templates"
var Files embed.FS

func main() {
	currentFolder, err := os.Getwd()
	if err != nil {
		log.Fatal("Something went wrong.: ", err)
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

	//creating component file

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

	//var tmplFiles = "templates"

	tmpl, err := template.ParseFS(Files, "./component.templ")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(file, componentTemplData)
	if err != nil {
		panic(err)
	}

	//creating component file

	//creating styles file

	file2, err := os.Create(path.Join(fullPath, "style.ts"))

	if err != nil {
		log.Fatal("File creation failed: ", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("File creation failed: ", err)
		}
	}(file2)

	componentTemplData2 := ComponentTemplData{
		Name: componentName,
	}

	//var tmplFiles = "templates"

	tmpl2, err := template.ParseFS(Files, "./style.templ")
	if err != nil {
		panic(err)
	}
	err = tmpl2.Execute(file2, componentTemplData2)
	if err != nil {
		panic(err)
	}

	//creating styles file

	//creating index file

	file3, err := os.Create(path.Join(fullPath, "index.ts"))

	if err != nil {
		log.Fatal("File creation failed: ", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("File creation failed: ", err)
		}
	}(file3)

	componentTemplData3 := ComponentTemplData{
		Name: componentName,
		File: folder,
	}

	//var tmplFiles = "templates"

	tmpl3, err := template.ParseFS(Files, "./index.templ")
	if err != nil {
		panic(err)
	}
	err = tmpl3.Execute(file3, componentTemplData3)
	if err != nil {
		panic(err)
	}

	//creating styles file
}
