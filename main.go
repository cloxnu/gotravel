package main

var (
	res = Res{}
	conf = Conf{}
)

func main() {
	res.load()
	conf.load()
	Compile()
}
