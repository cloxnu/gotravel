package main

import "fmt"

func main() {
	conf := Conf{}
	conf.load()
	fmt.Println(conf)
}
