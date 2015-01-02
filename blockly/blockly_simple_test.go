package blockly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadBlocklyFormat(t *testing.T) {
	testcase := `
	<xml xmlns="http://www.w3.org/1999/xhtml">
	  <block type="aidora_cluster" id="13" x="150" y="144">
	    <field name="NAME">name</field>
	    <field name="INSTANCES">1</field>
	    <statement name="pods">
	      <block type="aidora_pod" id="15">
	        <statement name="containers">
	          <block type="aidora_container" id="16" inline="false">
	            <field name="NAME">default</field>
	            <field name="INSTANCES">1</field>
	            <value name="RESOURCE">
	              <block type="aidora_container_resource" id="17">
	                <field name="CPU">50</field>
	                <field name="MEM">128</field>
	              </block>
	            </value>
	            <value name="IMAGE">
	              <block type="aidora_docker_image" id="18">
	                <field name="repository">repo</field>
	                <field name="name">name</field>
	                <field name="tag">latest</field>
	              </block>
	            </value>
	          </block>
	        </statement>
	      </block>
	    </statement>
	  </block>
	</xml>`

	conf := NewConfig(testcase)
	cluster := conf.Cluster(0)

	assert.Equal(t, "name", cluster.Name())
	assert.Equal(t, 1, cluster.Instances())

	pod := cluster.Pods(0)
	container := pod.Containers(0)

	assert.Equal(t, "default", container.Name())
	assert.Equal(t, 1, container.Instances())
	assert.Equal(t, 50, container.Cpu())
	assert.Equal(t, 128, container.Memory())
	assert.Equal(t, "repo/name:latest", container.Image())

}