package main

import (
	"embed"
	"gopkg.in/yaml.v2"
)

//go:embed res
var res embed.FS

type Res struct {
	BaseUrl string `yaml:"base_url"`
	Home map[string]string `yaml:"home"`
	Art map[string]string `yaml:"art"`
	Font map[string]string `yaml:"font"`
	Content string `yaml:"content"`
}

func (r *Res) load()  {
	file, err := res.ReadFile("res/res.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, r)
	if err != nil {
		panic(err)
	}
}
