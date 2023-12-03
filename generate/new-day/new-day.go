package main

import (
	"os"
	"text/template"
)

func main() {
	packageName := os.Args[1]

	partOneData := struct {
		PackageName string
		Part        string
	}{
		PackageName: packageName,
		Part:        "One",
	}

	partTwoData := struct {
		PackageName string
		Part        string
	}{
		PackageName: packageName,
		Part:        "Two",
	}

	templateFile, err := os.Open("new-day-template.go.tpl")
	if err != nil {
		panic(err)
	}
	defer templateFile.Close()

	b, err := os.ReadFile("new-day-template.go.tpl")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("template").Parse(string(b))
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("../../days/"+packageName, 0755)
	if err != nil {
		panic(err)
	}

	partOneFile, err := os.Create("../../days/" + packageName + "/part-one.go")
	if err != nil {
		panic(err)
	}
	defer partOneFile.Close()

	err = tmpl.ExecuteTemplate(partOneFile, "template", partOneData)
	if err != nil {
		panic(err)
	}

	partTwoFile, err := os.Create("../../days/" + packageName + "/part-two.go")
	if err != nil {
		panic(err)
	}
	defer partTwoFile.Close()

	err = tmpl.ExecuteTemplate(partTwoFile, "template", partTwoData)
	if err != nil {
		panic(err)
	}
}
