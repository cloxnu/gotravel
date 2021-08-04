package main

import (
	"io/ioutil"
	"net/url"
	"path"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	BaseUrl string `yaml:"base_url"`
	Top string `yaml:"top"`
	Content string `yaml:"content"`
}

func (c *Conf) load() {
	file, err := ioutil.ReadFile("./info.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}
}

func Url(p ...string) string {
	baseUrl := conf.BaseUrl
	u, err := url.Parse(baseUrl)
	if err != nil {
		panic(err)
	}

	p = append([]string{u.Path}, p...)
	u.Path = path.Join(p...)
	return u.String()
}

func StoryRelativeUrl(p ...string) string {
	if len(conf.BaseUrl) == 0 {
		p = append([]string{"../"}, p...)
	}
	return Url(p...)
}
