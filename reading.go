package gouiscreen

import (
	"github.com/goog-lukemc/gouidom"
	"github.com/goog-lukemc/gouielement"
)

type WorksData struct {
	Title    string
	Subtitle string
	Articles []*gouielement.ArticleData
}

// Reader create a readable content
func DocReader(parent string, comp *gouielement.ElementLib, works *WorksData) error {
	if works == nil {
		works = append(works, getDefaultContent(parent, v))
	}

	// Create a body centering div
	cd := comp.Div(parent, map[string]string{})

	// Create the left side bar which create a left had list of the documents in the collection
	// TODO: Implement
	comp.Div(gouidom.PathOf(cd), map[string]string{})

	for _, item := range *content {

	}
	// Content Body
	comp.Readable(content)

	// Create a right hand index of the sub headings in the document
	// TODO: Implement
	comp.Div(gouidom.PathOf(cd), map[string]string{})

}

func getDefaultContent(parent string, comp *gouielement.ElementLib) *WorksData {
	w := &WorksData{
		Title: "Something New!"
		Subtitle: "My Wife Made me do it."
		Articles: []&gouielement.ArticlesData{
			Title:"Welcome to the Go UI framework using Webassembly",
			Subtitle:"And Here we are!"
			Content:
				&gouielement.ContentCFG{
					Typ:gouidom.HTMLTag.Span,
					CFG:map[string]interface{}{
						"text":"The guiding principal of this project is to take a subset of the existing HTML, and javascript specifications and port them to gos webassembly implementation. From these building block UI can be represented in idiomatic go code without the need for markup."
					},
				}
			},
		},
	}
	

	return w
}
