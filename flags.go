package main

import "github.com/codegangsta/cli"

var (
	flDir = cli.StringFlag{
		Name:   "dir",
		Value:  "./static",
		Usage:  "Directory to serve static contents",
		EnvVar: "PRAKOB_DIR",
	}
)
