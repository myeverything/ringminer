/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package main

import(
	"github.com/Loopring/ringminer/cmd/utils"
	"sort"
	"gopkg.in/urfave/cli.v1"
	"os"
	"fmt"
	"github.com/Loopring/ringminer/node"
	"go.uber.org/zap"
	"github.com/Loopring/ringminer/config"
	"github.com/Loopring/ringminer/log"
	"os/signal"
)

var (
	app *cli.App
	configFile string
	globalConfig *config.GlobalConfig
	logger *zap.Logger
)

func main() {
	app = utils.NewApp()
	app.Action = minerNode
	app.HideVersion = true // we have a command to print the version
	app.Copyright = "Copyright 2013-2017 The Looprint Authors"
	app.Flags = []cli.Flag{cli.StringFlag{Name:"conf", Usage:" config file"}}


	app.Commands = []cli.Command{
		//matchengineCommand,
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(ctx *cli.Context) error {
		//runtime.GOMAXPROCS(runtime.NumCPU())
		file := ""
		if (ctx.IsSet(configFile)) {
			file = ctx.String("conf")
		}
		var err error
		globalConfig,err = config.LoadConfig(file)
		if nil != err {
			return err
		}


		logger = log.Initialize(globalConfig.LogOptions)
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer logger.Sync()
}

func minerNode(c *cli.Context) error {
	//todo：设置flag到config中
	n := node.NewNode(logger, globalConfig)
	n.Start()

	log.Info("started")
	//captiure stop signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)
	go func() {
		for {
			select {
			case sig := <- signalChan:
				log.Infof("captured %s, exiting...\n", sig.String())
				n.Stop()
				os.Exit(1)
			}
		}
	}()
	n.Wait()
	return nil
}
