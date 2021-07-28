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
	Top *Story
	Res Res
	Stories []Story
}

func Compile()  {
	copyResources()

	var err error
	tmpl, err = template.New("all_go_html").Funcs(template.FuncMap{
		"Url": func(path string) string { return conf.BaseUrl + path },
	}).ParseFS(templateFS, "template/*.gohtml")
	if err != nil {
		panic(err)
	}

	compileHome()
}

func copyResources()  {
	CopyRes("story")
	CopyRes("font")
	CopyRes("home")
	CopyRes("css")
	CopyRes("js")
}

func compileHome()  {
	file, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Funcs(template.FuncMap{
		"Url": func(path string) string { return conf.BaseUrl + path },
	}).ExecuteTemplate(file, "home.gohtml", CompileData{Conf: conf, Top: LoadStory(conf.Top), Res: res, Stories: stories})
	if err != nil {
		panic(err)
	}
}
