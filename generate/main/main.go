package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	data := struct {
		Days []string
	}{}
	b, err := os.ReadFile("main-template.go.tpl")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Parse(string(b))
	if err != nil {
		panic(err)
	}

	dir, err := os.Open("../../days")
	if err != nil {
		log.Fatal("Error opening directory:", err)
	}
	defer dir.Close()

	// Read the contents of the directory
	fileInfos, err := dir.ReadDir(0)
	if err != nil {
		log.Fatal("Error reading directory:", err)
	}

	dayFolders := []string{}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			dayFolders = append(dayFolders, fileInfo.Name())
		}
	}

	data.Days = make([]string, len(dayFolders))
	for i, dayFolder := range dayFolders {
		data.Days[i] = dayFolder
	}

	filePath := "../../main.go"

	if _, err := os.Stat(filePath); err == nil {
		os.Remove(filePath)
	}

	mainFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer mainFile.Close()

	err = tmpl.ExecuteTemplate(mainFile, "template", data)
	if err != nil {
		panic(err)
	}
}
