package main

import "github.com/codegangsta/cli"
import "github.com/zenazn/goji"
import "net/http"

func startServer(c *cli.Context) {
	dir := c.String("dir")
	goji.Handle("/*", http.FileServer(http.Dir(dir)))
	goji.Serve()
}
