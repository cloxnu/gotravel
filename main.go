package main

var (
	res     = Res{}
	conf    = Conf{}
	stories []Story
)

func main() {
	res.load()
	conf.load()
	stories = LoadStories()
	Compile()
}
