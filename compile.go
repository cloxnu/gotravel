package main

import (
	"embed"
	"html/template"
	"io/ioutil"
	"os"
	"path"
)

//go:embed template
var templateFS embed.FS

var tmpl *template.Template

type HomeData struct {
	Conf *Conf
	Top *Story
	Res *Res
	Stories []Story
}

type StoryData struct {
	Conf *Conf
	Res *Res
	Story *Story
	Content template.HTML
	AssociatedStories []Story
}

func Compile()  {
	copyResources()

	var err error
	tmpl, err = template.New("all_go_html").Funcs(template.FuncMap{
		"Url": Url,
		"mod": func(i, j int) bool { return i % j == 0 },
		"css": func(s string) template.CSS { return template.CSS(s) },
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
		"Url": Url,
		"mod": func(i, j int) bool { return i % j == 0 },
		"css": func(s string) template.CSS { return template.CSS(s) },
	}).ExecuteTemplate(file, "home.gohtml", HomeData{Conf: &conf, Top: LoadStory(conf.Top), Res: &res, Stories: stories})
	if err != nil {
		panic(err)
	}
}

func compileStories()  {
	for _, story := range stories {
		storyFile, err := ioutil.ReadFile(path.Join(conf.Content, story.Dir, story.Content))
		if err != nil {
			panic(err)
		}

		err = os.MkdirAll(story.Dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		outputHTML := Render(&story, storyFile)

		file, err := os.Create(path.Join(story.Dir, "index.html"))
		if err != nil {
			panic(err)
		}

		associatedStories := make([]Story, 0)
		for _, storyDir := range story.Associated {
			associatedStories = append(associatedStories, *LoadStory(storyDir))
		}

		err = tmpl.Funcs(template.FuncMap{
			"Url": StoryRelativeUrl,
			"mod": func(i, j int) bool { return i % j == 0 },
			"css": func(s string) template.CSS { return template.CSS(s) },
		}).ExecuteTemplate(file, "story.gohtml", StoryData{Conf: &conf, Res: &res, Story: &story, Content: template.HTML(outputHTML), AssociatedStories: associatedStories})
		if err != nil {
			panic(err)
		}

	}
}
