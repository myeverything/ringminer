package main

import(
	"github.com/Loopring/ringminer/cmd"
	"sort"
	"gopkg.in/urfave/cli.v1"
	"runtime"
	"os"
	"fmt"
)

var (
	app = cmd.NewApp()
)

// TODO(fukun): matchengine与order的通信

func init() {
	app.Action = miner
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2013-2017 The Looprint Authors"
	app.Commands = []cli.Command{

	}
	sort.Sort(cli.CommandsByName(app.Commands))

	//app.Flags = append(app.Flags, nodeFlags...)

	app.Before = func(ctx *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		return nil
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func miner() {

}

