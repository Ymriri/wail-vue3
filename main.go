package main

import (
	"embed"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"os"
	//"github.com/wailsapp/wails/v2/pkg/runtime"
	"path/filepath"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("get exe dir error!")
	}
	fmt.Println("workDir: " + dir)
	// 更改工作目录到程序的路径
	os.Chdir(dir)
	app := NewApp()
	// 加载配置文件
	app.LoadConfigInit()
	err = wails.Run(&options.App{
		Title:  "FtpScanner(87)",
		Width:  1524,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		Logger:             nil,
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		//BackgroundColour:   &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Windows: &windows.Options{
			Theme: windows.Light,
		},
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}

}
