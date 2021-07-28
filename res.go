package main

import (
	"embed"
	"gopkg.in/yaml.v2"
)

//go:embed res
var resFS embed.FS

type Res struct {
	Home map[string]string `yaml:"home"`
	Art map[string]string `yaml:"story"`
	Font map[string]string `yaml:"font"`
}

func (r *Res) load()  {
	file, err := resFS.ReadFile("res/res.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, r)
	if err != nil {
		panic(err)
	}
}
