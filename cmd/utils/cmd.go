package utils

import (
	"path/filepath"
	"os"
	"gopkg.in/urfave/cli.v1"
	"github.com/Loopring/ringminer/params"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = params.Version
	app.Usage = "the Loopring/ringminer command line interface"
	app.Author = ""
	app.Email = ""

	return app
}