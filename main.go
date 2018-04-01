package main

import (
	"os"

	"github.com/urfave/cli"
)

var Version string = "0.0.1"

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = "myhead"
	app.Usage = "myhead command"
	app.Version = Version
	app.Author = "yoichi watanabe"
	app.Action = doMyhead
	app.Flags = myHeadFlags
	return app
}
