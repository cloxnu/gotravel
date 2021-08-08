package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"sort"
)

type Story struct {
	Title string `yaml:"title"`
	Class string `yaml:"class"`
	Dir string `yaml:"dir"`
	Cover string `yaml:"cover"`
	Description string `yaml:"description"`
	Content string `yaml:"content"`
	Priority int64 `yaml:"priority"`
	StoryDate string `yaml:"story_date"`
	CreationDate string `yaml:"creation_date"`
	ModificationDate string `yaml:"modification_date"`
	Associated []string `yaml:"associated"`
}

func (s *Story) Path(p ...string) string {
	p = append([]string{conf.Content, s.Dir}, p...)
	return Url(p...)
}

func (s *Story) CoverPath() string {
	if len(s.Cover) == 0 {
		return ""
	}
	return s.Path(s.Cover)
}

func (s *Story) StoryRelativePath(p ...string) string {
	p = append([]string{conf.Content, s.Dir}, p...)
	return StoryRelativeUrl(p...)
}

func (s *Story) StoryRelativeCoverPath() string {
	if len(s.Cover) == 0 {
		return ""
	}
	return s.StoryRelativePath(s.Cover)
}

func (s *Story) ClassificationColor() string {
	switch s.Class {
	case "travel": return "--orange"
	case "life": return "--teal"
	case "inspiration": return "--indigo"
	default: return "--foreground-color"
	}
}

func LoadStories() []Story {
	stories := make([]Story, 0)
	
	storyDir, err := ioutil.ReadDir(conf.Content)
	if err != nil {
		panic(err)
	}

	for _, dir := range storyDir {
		configFilePath := path.Join(conf.Content, dir.Name(), "/info.yaml")
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

	sort.Slice(stories, func(i, j int) bool {
		return stories[i].Priority > stories[j].Priority
	})

	return stories
}

func LoadStory(dir string) *Story {
	configFilePath := path.Join(conf.Content, dir, "info.yaml")
	if IsFileExist(configFilePath) {
		story := &Story{}
		configFile, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(configFile, story)
		if err != nil {
			panic(err)
		}
		return story
	}
	return nil
}
