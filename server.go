package main

import "github.com/codegangsta/cli"
import "github.com/zenazn/goji"
import "github.com/zenazn/goji/web"
import "net/http"
import "fmt"
import "github.com/aidora/prakob/blockly"
import "io/ioutil"

func startServer(c *cli.Context) {
	goji.Get("/load/:name", func(c web.C, w http.ResponseWriter, r *http.Request) {
		filename := "./db/" + c.URLParams["name"] + ".xml"
		bytes, _ := ioutil.ReadFile(filename)
		fmt.Fprint(w, string(bytes))
	})
	goji.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		content := r.Form["xml"][0]
		filename := r.Form["workspace"][0]
		_ = blockly.NewConfig(content)
		ioutil.WriteFile("./db/" + filename + ".xml", []byte(content), 0644)
		http.Error(w, http.StatusText(200), 200)
	})

	dir := c.String("dir")
	goji.Handle("/*", http.FileServer(http.Dir(dir)))
	goji.Serve()
}
