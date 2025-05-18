package view

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"
)

func InlineSVG(filename string, className ...string) template.HTML {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("could not open svg %s: %v", filename, err)
	}

	svg := string(content)

	if len(className) > 0 {
		classes := strings.Join(className, " ")
		svg = appendSvgClass(svg, classes)
	}

	return template.HTML(svg)
}

type svgType struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
}

// Appends classes to supplied SVG element. Ensures to keep any existing
// classes if they exist on the SVG already.
func appendSvgClass(svgStr string, classes string) string {
	decoder := xml.NewDecoder(strings.NewReader(svgStr))

	// Split the SVG into its attributes
	var svg svgType
	err := decoder.Decode(&svg)
	if err != nil {
		log.Fatal("error decoding SVG: %w", err)
	}

	// Append or add the additional classes
	classFound := false
	for i, attr := range svg.Attrs {
		if attr.Name.Local == "class" {
			existing := strings.Fields(attr.Value)
			existing = append(existing, classes)
			svg.Attrs[i].Value = strings.Join(existing, " ")
			classFound = true
			break
		}
	}
	if !classFound {
		svg.Attrs = append(svg.Attrs, xml.Attr{Name: xml.Name{Local: "class"}, Value: classes})
	}

	// Put SVG back together with updated attribute values
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("<%s", svg.XMLName.Local))
	for _, attr := range svg.Attrs {
		buf.WriteString(fmt.Sprintf(` %s="%s"`, attr.Name.Local, attr.Value))
	}
	buf.WriteString(">")
	buf.WriteString(string(svg.Content))
	buf.WriteString(fmt.Sprintf("</%s>", svg.XMLName.Local))

	return buf.String()
}
