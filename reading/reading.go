package gouiscreen

import (
	"gihthub.com/goog-lukemc/gouielement"
)

func Reader(content *gouielement.ReadingData, comp *gouielement.ElementLib, parent string) error {
	if content == nil {
		content = getDefaultContent(parent, v)
	}
	comp.Readable(content)

}

func getDefaultContent(parent string, comp *gouielement.ElementLib) *gouielement.ReadingData {
	c := &gouielement.ReadingData{
		Title:    "Welcome to the Go UI framework using Webassembly",
		Subtitle: "Calm down!, with project is a WIP",
	}
	t1 := "Welcome to the goui framework. Have fun!"
	c.Content = append(c.Content,
		comp.Span(parent,t1)
	)

	return c
}
