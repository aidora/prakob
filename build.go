package main

import "github.com/codegangsta/cli"
import "fmt"
import yaml "gopkg.in/yaml.v2"
import "io/ioutil"

type Node struct {
	Image   string
	Scale   int
	Links   []string
	Build   string
	Volumns []string
	Net     string
	Ports   []string
}

type Pod struct {
	Type  string
	Nodes []map[string]Node
}

type Config struct {
	Cluster struct {
		Scale int
		Pods  []Pod
	}
}

func build(c *cli.Context) {
	data, err := ioutil.ReadFile("./prakob.yml")
	fmt.Printf("%s\n----------\n", string(data))
	config := &Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", config.Cluster.Pods)
}
