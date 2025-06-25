package striphtml

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/net/html"
)

func StripHTML(input string) (string, error) {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %w", err)
	}

	var buf bytes.Buffer
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(doc)

	// Replace non-breaking spaces with normal spaces
	text := strings.ReplaceAll(buf.String(), "\u00A0", " ")

	// Optionally trim and normalize whitespace
	out := strings.Map(func(r rune) rune {
		if unicode.IsControl(r) && r != '\n' && r != '\r' {
			return -1
		}
		return r
	}, text)

	return strings.TrimSpace(out), nil
}
