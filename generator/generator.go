package generator

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type FileMetaData struct {
	FolderName string
	Name       string
	FolderPath string
}

//go:embed templates
var Files embed.FS

func CreateFile(metaData FileMetaData, filename, templateName string) error {
	fullPath := filepath.Join(metaData.FolderPath, filename)

	// Ensure the parent directory exists
	if err := os.MkdirAll(metaData.FolderPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Panic("File creation failed: ", err)
		}
	}(file)

	tmpl, err := template.ParseFS(Files, fmt.Sprintf("templates/%s.templ", templateName))
	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, metaData); err != nil {
		return err
	}

	return nil
}

func GetMetaData(packageType string, packageName string) FileMetaData {
	currentFolder, err := os.Getwd()
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}

	if len(packageName) < 3 {
		log.Fatal("Folder name is too short")
	}

	fullPath := filepath.Join(currentFolder, packageType, packageName)

	if err != nil {
		log.Fatal("Folder creation failed: ", err)
	}

	names := strings.Split(packageName, "-")
	for i, name := range names {
		newName := strings.ToUpper(name[:1]) + strings.ToLower(name[1:])
		names[i] = newName
	}

	componentName := strings.Join(names, "")

	return FileMetaData{
		Name:       componentName,
		FolderPath: fullPath,
		FolderName: packageName,
	}
}
