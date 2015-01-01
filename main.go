package main

import "github.com/codegangsta/cli"
import "os"
import "log"

func main() {
	app := cli.NewApp()
	app.Name = "prakob"
	app.Usage = "docker cluster builder"
	app.Version = "0.0.1"
	app.Author = ""
	app.Email = ""

	app.Commands = []cli.Command{
		{
			Name:      "build",
			ShortName: "b",
			Usage:     "build cluster",
			Action:    build,
		},
		{
			Name:      "start",
			ShortName: "s",
			Usage:     "start server",
			Action:    startServer,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
