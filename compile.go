package main

import (
	"embed"
	"html/template"
	"os"
)

//go:embed template
var templateFS embed.FS

var tmpl *template.Template

type CompileData struct {
	Conf Conf
	Res Res
}

func Compile()  {
	var err error
	tmpl, err = template.ParseFS(templateFS, "template/*.gohtml")
	if err != nil {
		panic(err)
	}

	compileLoading()
}

func compileLoading()  {
	file, err := os.Create("loading.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(file, "loading.gohtml", CompileData{Conf: conf, Res: res})
	if err != nil {
		panic(err)
	}
}
