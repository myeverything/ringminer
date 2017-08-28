package cmd

import (
	"path/filepath"
	"os"
	"gopkg.in/urfave/cli.v1"
	"github.com/Loopring/ringminer/config"
)

func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = config.Version
	app.Usage = "the Loopring/ringminer command line interface"
	app.Author = ""
	app.Email = ""

	return app
}