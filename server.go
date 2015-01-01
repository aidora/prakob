package main

import "github.com/codegangsta/cli"
import "github.com/zenazn/goji"
import "net/http"

func startServer(c *cli.Context) {
    goji.Handle("/*", http.FileServer(http.Dir("../static")))
    goji.Serve()
}