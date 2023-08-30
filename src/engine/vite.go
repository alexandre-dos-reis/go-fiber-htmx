package engine

import (
	"github.com/gofiber/template/html/v2"
)

func AddViteAssets(engine *html.Engine, isDev bool) {
	engine.AddFunc("vite_assets", func() string {
		if isDev {
			return `
      <script type="module" src="/@vite/client"></script>
      <script type="module" src="http://localhost:5173/main.ts" defer></script>
      `
		} else {
			return `
      <link rel="stylesheet" href="/assets/main.css">
      <script type="module" src="/assets/main.js" defer></script>
      `
		}
	})
}
