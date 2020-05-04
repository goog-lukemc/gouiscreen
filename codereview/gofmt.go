package codereview

import (
	"go/scanner"
	"go/token"
	"strings"
	"unicode/utf8"
)

type codeDetail struct {
	name         string
	start        int
	token        *token.Token
	value        string
	lit          string
	isKeyWord    bool
	isOperator   bool
	isLiteral    bool
	isExported   bool
	isIdentifier bool
}

func getCodeTokens(name string, bts []byte) []*codeDetail {

	// Clean the string just in case of special charters
	line := getStringFromBytes(bts)

	// Create out results object
	tres := []*codeDetail{}
	ares := []*codeDetail{}

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()
	fs := fset.AddFile(name, fset.Base(), len(bts))

	// Init the scanner with no error handler.
	s.Init(fs, bts, nil, scanner.ScanComments)

	// Scan the file lines
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		p := fset.Position(pos)

		// Trying this.. Delete maybe
		if lit == "" {
			lit = tok.String()
		}
		///////////////////////////
		tres = append(tres, &codeDetail{
			name:         tok.String(),
			start:        p.Offset,
			token:        &tok,
			value:        lit,
			isLiteral:    tok.IsLiteral(),
			isOperator:   tok.IsOperator(),
			isIdentifier: token.IsIdentifier(lit),
			isExported:   token.IsExported(lit),
			isKeyWord:    tok.IsKeyword(),
			lit:          lit,
		})

	}

	for pos, res := range tres {

		// No token in position 0. App a token for that position
		if pos == 0 {
			if res.start != 0 {
				// Add that string
				ares = append(ares, newCodeDetails(nil, line, 0, res.start))
			}
			// It is first position so add it.

			ares = append(ares, res)
			continue
		}

		// If the previous end != the start - add the gap
		lastEnd := (tres[pos-1].start + len(tres[pos-1].value))
		if lastEnd < res.start {
			ares = append(ares, newCodeDetails(nil, line, lastEnd, res.start))
		}

		// Check is this is the last iteration. If yes then add the remaining charters.
		if pos+1 == len(tres) {
			if (len(res.value) + res.start) == len(line) {
				if res.token != nil {
					ares = append(ares, res)
				}
				continue
			}
			if res.start != len(line) {
				ares = append(ares, newCodeDetails(res.token, line, res.start, (len(line)-1)))
			}
		}

		ares = append(ares, res)
	}

	return ares
}

func newCodeDetails(tok *token.Token, line string, start, end int) *codeDetail {
	segment := getLineSegment(line, start, end)
	if segment == "" {
		if tok != nil {
			segment = tok.String()
		}
	}

	return &codeDetail{
		start: start,
		token: tok,
		value: segment,
	}

}

func getLineSegment(line string, start int, end int) string {
	s := line[start:end]
	return s
}

func getStringFromBytes(bts []byte) string {

	sb := strings.Builder{}

	for len(bts) > 0 {
		r, size := utf8.DecodeRune(bts)
		sb.WriteRune(r)
		bts = bts[size:]
	}
	return sb.String()
}
