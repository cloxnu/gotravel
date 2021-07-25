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
	Stories []Story
}

func Compile()  {
	createOutputDirectory()
	copyResources()

	var err error
	tmpl, err = template.ParseFS(templateFS, "template/*.gohtml")
	if err != nil {
		panic(err)
	}

	compileHome()
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
	CopyRes("css")
	CopyRes("js")
}

func compileHome()  {
	file, err := os.Create(conf.Out + "index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(file, "home.gohtml", CompileData{Conf: conf, Res: res, Stories: stories})
	if err != nil {
		panic(err)
	}
}
