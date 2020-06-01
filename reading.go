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
func DocReader(parent string, comp *gouielement.ElementLib, works *WorksData) {
	if works == nil {
		works = getDefaultContent()
	}

	// Create a body centering div
	cd := comp.Div(parent, map[string]string{})

	// Create the left side bar which create a left had list of the documents in the collection

	wk := comp.Div(gouielement.PathOf(cd), map[string]string{}, "works-index")
	comp.UnOrderedList(gouielement.PathOf(wk), getWorksIndex(works))

	// Content Body
	for _, a := range works.Articles {
		comp.Readable(gouielement.PathOf(cd), a)
	}

	// Create a right hand index of the sub headings in the document
	// TODO: Implement
	comp.Div(gouielement.PathOf(cd), map[string]string{}, "article-index")

}

func getWorksIndex(w *WorksData) []string {
	at := []string{}
	for _, item := range w.Articles {
		at = append(at, item.Title)
	}
	return at
}

func getDefaultContent() *WorksData {
	w := &WorksData{
		Title:    "Something New!",
		Subtitle: "My Wife Made me do it.",
		Articles: []*gouielement.ArticleData{
			&gouielement.ArticleData{
				Title:    "Welcome to the Go UI framework using Webassembly",
				Subtitle: "And Here we are!",
				Content: []*gouielement.ContentConfig{
					&gouielement.ContentConfig{
						Typ: gouidom.HTMLTag.Img,
						CFG: map[string]interface{}{
							"ea": map[string]string{"src": "/image/go-logo_blue.svg"},
						},
					},
					&gouielement.ContentConfig{
						Typ: gouidom.HTMLTag.Span,
						CFG: map[string]interface{}{
							"secheading": "The Good Stuff",
							"text":       "The guiding principal of this project is to take a subset of the existing HTML, and javascript specifications and port them to gos webassembly implementation. From these building block UI can be represented in idiomatic go code without the need for markup.",
						},
					},
				},
			},
		},
	}

	return w
}
