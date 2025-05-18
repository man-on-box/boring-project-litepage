package view

import (
	"encoding/xml"
	"html/template"
	"log"
	"os"
)

func InlineSVG(filename string, className ...string) template.HTML {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not open svg %s: %v", filename, err)
	}

	svg := string(content)

	return template.HTML(svg)
}

type svgType struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
}
