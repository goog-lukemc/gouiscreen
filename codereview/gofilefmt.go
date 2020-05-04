package codereview

import (
	"bufio"
	"fmt"
	"gouicomponent"
	"gouidom"
	"io"
	"path"
	"strconv"
)

func GoFMTWeb(parent string, file io.Reader, v *gouidom.VDOM) {
	// Create a new component lib
	comp := gouicomponent.NewComponentLib(v)

	// Create a go formater
	//var parser wasmGoFMT

	container := comp.NewDivWrapper(parent)

	// Create a wrapper for the item
	//container := gouicomponent.NewWrapper(parent, v)

	lineScanner := bufio.NewScanner(file)

	// Split the file on lines
	lineScanner.Split(bufio.ScanLines)

	// Loop and generate the document
	l := 1
	for lineScanner.Scan() {

		// Line number parent div
		parentSpan := comp.NewDiv(path.Join(container.Parent, container.ID), map[string]string{"data-num": fmt.Sprintf("%d", l)}, "golang", "line")

		// Line Span
		lineToolBar := comp.NewSpan(path.Join(parentSpan.Parent, parentSpan.ID), "", "golang", "line-toolbar")

		// Line Span
		_ = comp.NewSpan(path.Join(lineToolBar.Parent, lineToolBar.ID), fmt.Sprintf("%d", l), "golang", "line-num")

		// Get the text from the reader
		text := lineScanner.Text()

		// Get the line results and loop to build the spans
		codeScan := getCodeTokens("unknown", []byte(text))

		// Loop the code scan line and set classes
		for pos, lr := range codeScan {

			if lr.value == "" {
				lr.value = "->"
			}

			_ = comp.NewSpan(path.Join(parentSpan.Parent, parentSpan.ID), lr.value, getSegmentClass(lr, codeScan, pos)...)

		}

		l++
	}
}

func getSegmentClass(seg *codeDetail, res []*codeDetail, pos int) []string {
	c := []string{"line-segment"}

	switch {
	case seg.isKeyWord:
		c = append(c, "keyword")
	case seg.name == "STRING":
		c = append(c, "text-string")
	case seg.isExported && seg.name == "IDENT":

		if pos != 0 {
			if res[pos-1].lit == "." {
				if res[pos+1].lit == "(" {
					c = append(c, "function")
				} else {
					c = append(c, "property")
				}
			}
		}
	case seg.name == "COMMENT":
		c = append(c, "comment")
	}

	return c
}

// HighlightLines raises the code line to the users attention and returns the center vertical position
// of the highlighted block.
func HighlightLines(lines []string) string {
	for pos, lineNumber := range lines {
		ele, err := gouidom.GetElementByID(lineNumber)
		if err != nil {
			gouidom.CLog("%s", err.Error())
		}
		if pos == 0 {
			ele.ScrollIntoView(true)
			ele.AddClass("top-shadow")
		}
		ele.AddClass("highlight")
		if pos == len(lines)-1 {
			ele.AddClass("bottom-shadow")
		}
	}
	return strconv.Itoa(len(lines) / 2)
}

func ScrollPositionCenter(center string) {
	// Get the window size with window

}
