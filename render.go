package main

import (
	"github.com/russross/blackfriday/v2"
	"io"
	"path"
)

type HTMLRenderer struct {
	UrlProcessor func([]byte) []byte
	*blackfriday.HTMLRenderer
}

func (r *HTMLRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if node.Type == blackfriday.Image {
		node.LinkData.Destination = r.UrlProcessor(node.LinkData.Destination)
	}
	return r.HTMLRenderer.RenderNode(w, node, entering)
}

func Render(story *Story, content []byte) []byte {
	r := &HTMLRenderer{
		UrlProcessor: func(input []byte) []byte {
			return []byte(path.Join(story.Path(), string(input)))
		},
		HTMLRenderer: blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			HeadingIDPrefix: "title-anchor-",
		}),
	}
	return blackfriday.Run(content, blackfriday.WithRenderer(r), blackfriday.WithExtensions(blackfriday.AutoHeadingIDs))
}

