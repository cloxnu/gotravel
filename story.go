package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Story struct {
	Title string `yaml:"title"`
	Class string `yaml:"class"`
	Dir string `yaml:"dir"`
	Cover string `yaml:"cover"`
	Description string `yaml:"description"`
	Content string `yaml:"content"`
	StoryDate string `yaml:"story_date"`
	CreationDate string `yaml:"creation_date"`
	ModificationDate string `yaml:"modification_date"`
	Associated []string `yaml:"associated"`
}

func LoadStories() []Story {
	stories := make([]Story, 0)
	
	storyDir, err := ioutil.ReadDir(conf.Content)
	if err != nil {
		panic(err)
	}

	for _, dir := range storyDir {
		configFilePath := conf.Content + dir.Name() + "/info.yaml"
		if dir.IsDir() && IsFileExist(configFilePath) {
			story := &Story{}
			configFile, err := ioutil.ReadFile(configFilePath)
			if err != nil {
				panic(err)
			}
			err = yaml.Unmarshal(configFile, story)
			if err != nil {
				panic(err)
			}
			stories = append(stories, *story)
		}
	}

	return stories
}
