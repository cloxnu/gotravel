package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	BaseUrl string `yaml:"base_url"`
	Top string `yaml:"top"`
	Content string `yaml:"content"`
	Out string `yaml:"out"`
}

func (c *Conf) load()  {
	file, err := ioutil.ReadFile("./info.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}
}
