package main

import "github.com/codegangsta/cli"
import "github.com/zenazn/goji"
import "net/http"
import "fmt"

func startServer(c *cli.Context) {
	goji.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form["xml"])
		http.Error(w, http.StatusText(200), 200)
	})

	dir := c.String("dir")
	goji.Handle("/*", http.FileServer(http.Dir(dir)))
	goji.Serve()
}
