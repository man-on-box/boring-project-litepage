package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"time"

	"github.com/man-on-box/boring-project-litepage/view"
	"github.com/man-on-box/litepage"
)

func main() {
	lp, err := litepage.New("man-on-box.github.io", litepage.WithBasePath("/boring-project-litepage"))
	if err != nil {
		log.Fatalf("Could not create app: %v", err)
	}

	lp.Page("/index.html", handleHomepage())

	err = lp.BuildOrServe()
	if err != nil {
		log.Fatalf("Could not run app: %v", err)
	}
}

var tmpl = template.New("").Funcs(template.FuncMap{
	"version": func() string {
		return fmt.Sprintf("%d", time.Now().Unix())
	},
	"inlineSVG": view.InlineSVG,
})

type Metric struct {
	Value string
	Label string
	Icon  string
}

func handleHomepage() func(w io.Writer) {
	data := struct {
		Metrics []Metric
	}{}

	data.Metrics = []Metric{
		{Value: "174", Label: "Lines of code", Icon: "assets/icons/code.svg"},
		{Value: "2", Label: "Dependencies", Icon: "assets/icons/cube.svg"},
		{Value: "0", Label: "Dependabot alerts", Icon: "assets/icons/alert.svg"},
		{Value: "100", Label: "Performance", Icon: "assets/icons/chart.svg"},
	}

	t := template.Must(tmpl.ParseFiles("./view/layouts/base.html", "./view/index.html"))

	return func(w io.Writer) {
		err := t.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Printf("error while rendering template: %v", err)
		}
	}

}
