package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

func CreateComponentFile(metaData FileMetaData, filename, templateName string) error {
	file, err := os.Create(path.Join(metaData.ComponentFolderPath, filename))

	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Panic("File creation failed: ", err)
		}
	}(file)

	tmpl, err := template.ParseFS(Files, fmt.Sprintf("templates/%s.templ", templateName))
	if err != nil {
		return err
	}
	err = tmpl.Execute(file, metaData)
	if err != nil {
		return err
	}

	return nil
}
