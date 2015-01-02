package blockly

import "strconv"
import "encoding/xml"

type Block struct {
	XmlName xml.Name     `xml:"block"`
	Type    string       `xml:"type,attr"`
	Field   []*Field     `xml:"field"`
	Stmt    []*Statement `xml:"statement"`
	Value   []*Value     `xml:"value"`
	Next    *Block       `xml:"next>block"`
}

func (b *Block) checkType(typeName string) {
	if b.Type != typeName {
		panic("Type not correct: " + b.Type)
	}
}

func (b *Block) getField(name string) string {
	for _, f := range b.Field {
		if f.Name == name {
			return f.Value
		}
	}
	return ""
}

func (b *Block) getValue(name string) *Value {
	for _, v := range b.Value {
		if v.Name == name {
			return v
		}
	}
	return nil
}

type Field struct {
	XmlName xml.Name `xml:"field"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",innerxml"`
}

type Statement struct {
	XmlName xml.Name `xml:"statement"`
	Name    string   `xml:"name,attr"`
	Block   *Block   `xml:"block"`
}

type Value struct {
	XmlName xml.Name `xml:"value"`
	Name    string   `xml:"name,attr"`
	Block   *Block   `xml:"block"`
}

// =========

type Config struct {
	XmlName xml.Name `xml:"xml"`
	Block   []*Block `xml:"block"`
}

func (c *Config) Cluster(i int) *Cluster {
	return (*Cluster)(c.Block[i])
}

func NewConfig(input string) *Config {
	conf := &Config{}
	xml.Unmarshal([]byte(input), conf)
	return conf
}

// =========

type Cluster Block
func (c *Cluster) Name() string {
	b := (*Block)(c)
	b.checkType("aidora_cluster")
	return b.getField("NAME")
}

func (c *Cluster) Instances() (i int) {
	b := (*Block)(c)
	b.checkType("aidora_cluster")
	i, _ = strconv.Atoi(b.getField("INSTANCES"))
	return
}

func (c *Cluster) Pods(i int) *Pod {
	(*Block)(c).checkType("aidora_cluster")

	if c.Stmt[0].Name != "pods" {
		panic("Not Pod")
	}
	if i == 0 {
		return (*Pod)(c.Stmt[0].Block)
	}
	return nil
}

type Pod Block
func (p *Pod) Containers(i int) *Container {
	(*Block)(p).checkType("aidora_pod")
	return (*Container)(p.Stmt[0].Block)
}

type Container Block
func (c *Container) Name() string {
	b := (*Block)(c)
	b.checkType("aidora_container")
	return b.getField("NAME")
}

func (c *Container) Instances() (i int) {
	b := (*Block)(c)
	b.checkType("aidora_container")
	i, _ = strconv.Atoi(b.getField("INSTANCES"))
	return
}

func (c *Container) Cpu() (i int) {
	b := (*Block)(c)
	b.checkType("aidora_container")
	r := b.getValue("RESOURCE")
	i, _ = strconv.Atoi(r.Block.getField("CPU"))
	return
}

func (c *Container) Memory() (i int) {
	b := (*Block)(c)
	b.checkType("aidora_container")
	r := b.getValue("RESOURCE")
	i, _ = strconv.Atoi(r.Block.getField("MEM"))
	return
}

func (c *Container) Image() string {
	b := (*Block)(c)
	b.checkType("aidora_container")
	r := b.getValue("IMAGE")
	rb := r.Block
	return rb.getField("repository") + "/" + rb.getField("name") + ":" + rb.getField("tag")
}
