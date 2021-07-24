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
	createOutputDirectory()
	copyResources()

	var err error
	tmpl, err = template.ParseFS(templateFS, "template/*.gohtml")
	if err != nil {
		panic(err)
	}

	//compileLoading()
}

func createOutputDirectory() {
	err := os.Mkdir(conf.Out, os.ModePerm)
	if err != nil {
		return
	}
}

func copyResources()  {
	CopyRes("art")
	CopyRes("font")
	CopyRes("home")
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
