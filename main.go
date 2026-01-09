package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed assets/dict/s_1.txt
var dictEmbed embed.FS

//go:embed assets/input/input.txt
var inputEmbed string

func main() {
	// Create an instance of the app structure
	searchEngine := newSearchEngine()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "App",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        searchEngine.startup,
		Bind: []interface{}{
			searchEngine,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
