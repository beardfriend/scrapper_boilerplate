package main

import (
	"boilerplate/internal/modules/meta"
	"boilerplate/internal/wire"
	"context"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	meta := &meta.Meta{
		Name:        "보일러 플레이트",
		Description: "보일러 플레이트 설명",
		VersionText: "V1.0.0",
		VersionSum:  10000,
		CreatedAt:   "2024-02-19T12:43:00",
	}
	app := wire.InitializeSampleApp(meta)

	err := wails.Run(&options.App{
		Title:           meta.Name,
		Width:           1440,
		Height:          768,
		CSSDragProperty: "--wails-draggable",
		CSSDragValue:    "drag",
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup: func(ctx context.Context) {
			app.StartUp(ctx)
		},
		Bind: app.GetHandlers(),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
