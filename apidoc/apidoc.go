package apidoc

import (
	"gouicomponent"
	"gouidom"
)

type Doc struct {
	Name    string
	Content *Article
	Demo    *HowTo
}

type Article struct {
	Heading    string
	SubHeading string
	Sections   []*Section
}

type Section struct {
	Heading    string
	Paragraphs []string
}

type HowTo struct {
	Method       string
	RequestBody  map[string]interface{}
	ResponseBody map[string]interface{}
	Path         string
	Example      string
}

func NewAPIDoc(v *gouidom.VDOM, d []*Doc) []*gouidom.Element {

	// Create a new component lib
	comp := gouicomponent.NewComponentLib(v)

	// Create our main content with 3 divs
	content := comp.NewGrid("html/body", 3, "container", "grid")

	_ = content

	names := []string{}
	for _, doc := range d {
		names = append(names, doc.Name)
	}

	comp.NewButtonGroup(gouicomponent.PathOf(content[0]), "Routes", names...)

	// Lets build a document
	c1 := comp.NewArticle(gouicomponent.PathOf(content[1]), "article-outer")

	for _, doc := range d {
		comp.NewHeading(gouicomponent.PathOf(c1), doc.Content.Heading, "1", "heading-main")
		for _, sec := range doc.Content.Sections {
			s := comp.NewSection(gouicomponent.PathOf(c1), "section-inner")
			comp.NewHeading(gouicomponent.PathOf(s), sec.Heading, "3", "heading-section")
			for _, para := range sec.Paragraphs {
				comp.NewParagraph(gouicomponent.PathOf(s), para, "para-text")
			}
		}
		r1 := comp.NewArticle(gouicomponent.PathOf(content[2]), "article-example")
		s1 := comp.NewSection(gouicomponent.PathOf(r1))
		comp.NewPreCode(gouicomponent.PathOf(s1), doc.Demo.Example)
	}

	return content

}
