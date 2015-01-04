package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "net/http"
import "strings"

import "github.com/aidora/prakob/blockly"
import "github.com/codegangsta/cli"
import "github.com/zenazn/goji"
import "github.com/zenazn/goji/web"

func startServer(c *cli.Context) {
	goji.Post("/deploy/:name", func(c web.C, w http.ResponseWriter, r *http.Request) {
		filename := c.URLParams["name"]
		bytes, _ := ioutil.ReadFile("./db/" + filename)
		conf := blockly.NewConfig(string(bytes))
		clusterName := conf.Cluster(0).Name()
		fmt.Fprint(w, "Deploying: "+clusterName)
	})

	goji.Get("/filenames", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		files, _ := ioutil.ReadDir("./db")
		result := make(map[string][]string)
		list := make([]string, 0)
		for _, f := range files {
			list = append(list, f.Name())
		}
		result["filenames"] = list
		output, _ := json.Marshal(result)
		fmt.Fprint(w, string(output))
	})

	goji.Get("/load/:name", func(c web.C, w http.ResponseWriter, r *http.Request) {
		filename := "./db/" + c.URLParams["name"]
		bytes, _ := ioutil.ReadFile(filename)
		fmt.Fprint(w, string(bytes))
	})

	goji.Post("/save", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		content := r.Form["xml"][0]
		filename := r.Form["workspace"][0]
		if strings.HasSuffix(filename, ".xml") == false {
			filename = filename + ".xml"
		}
		ioutil.WriteFile("./db/"+filename, []byte(content), 0644)
		http.Error(w, http.StatusText(200), 200)
	})

	dir := c.String("dir")
	goji.Handle("/*", http.FileServer(http.Dir(dir)))
	goji.Serve()
}
