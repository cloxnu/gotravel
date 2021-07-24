package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Version string `yaml:"version"`
	Top string `yaml:"top"`
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
