package gendoc

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

var (
	paraPattern         = regexp.MustCompile(`(\n|\r|\r\n)\s*`)
	paraPattern2        = regexp.MustCompile(`(\n|\r|\r\n)`)
	spacePattern        = regexp.MustCompile("( )+")
	multiNewlinePattern = regexp.MustCompile(`(\r\n|\r|\n){2,}`)
)

// PFilter splits the content by new lines and wraps each one in a <p> tag.
func PFilter(content string) template.HTML {
	paragraphs := paraPattern.Split(content, -1)
	return template.HTML(fmt.Sprintf("<p>%s</p>", strings.Join(paragraphs, "</p><p>")))
}

// ParaFilter splits the content by new lines and wraps each one in a <para> tag.
func ParaFilter(content string) string {
	paragraphs := paraPattern.Split(content, -1)
	return fmt.Sprintf("<para>%s</para>", strings.Join(paragraphs, "</para><para>"))
}

// MDBrFilter splites the content by new lines and postfix each one with <br/>.
func MDBrFilter(content string) string {
	paragraphs := paraPattern.Split(content, -1)
	return strings.Join(paragraphs, "<br/>")
}

// MDBlockFilter splites the content by new lines and prefix each one with \t except start with ^
func MDBlockFilter(content string) string {
	paragraphs := paraPattern2.Split(content, -1)
	res := "\t" + strings.Join(paragraphs, "\n\t")
	return strings.Replace(res, "\t^", "", -1)
}

// NoBrFilter removes single CR and LF from content.
func NoBrFilter(content string) string {
	normalized := strings.Replace(content, "\r\n", "\n", -1)
	paragraphs := multiNewlinePattern.Split(normalized, -1)
	for i, p := range paragraphs {
		withoutCR := strings.Replace(p, "\r", " ", -1)
		withoutLF := strings.Replace(withoutCR, "\n", " ", -1)
		paragraphs[i] = spacePattern.ReplaceAllString(withoutLF, " ")
	}
	return strings.Join(paragraphs, "\n\n")
}
