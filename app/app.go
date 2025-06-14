package app

import "github.com/urfave/cli"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Search for IPs and Servers name in web"

	return app
}
