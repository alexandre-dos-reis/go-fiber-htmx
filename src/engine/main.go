package engine

import (
	"html/template"

	"github.com/gofiber/template/html/v2"
)

func GetEngine() *html.Engine {
	engine := html.New("./src/views", ".html")
	AddViteAssets(engine, false)
	addUnescape(engine)

	return engine
}

func addUnescape(engine *html.Engine) {
	engine.AddFunc(
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)
}
