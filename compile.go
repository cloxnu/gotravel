package main

import (
	"embed"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io/ioutil"
	"os"
	"path"
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
	compileStories()
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

func compileStories()  {
	for _, story := range stories {
		storyFile, err := ioutil.ReadFile(path.Join(story.Path(), story.Content))
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(story.Dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		blackfriday.HTMLRenderer{}
		outputHTML := blackfriday.Run(storyFile)
		err = ioutil.WriteFile(path.Join(story.Dir, "index.html"), outputHTML, os.ModePerm)
		if err != nil {
			panic(err)
		}

	}
}
