package blockly

var complex = `
<xml xmlns="http://www.w3.org/1999/xhtml">
  <block type="aidora_cluster" id="13" x="199" y="-145">
    <field name="NAME">name</field>
    <field name="INSTANCES">1</field>
    <statement name="pods">
      <block type="aidora_pod" id="14">
        <statement name="containers">
          <block type="aidora_container" id="16" inline="false">
            <field name="NAME">default</field>
            <field name="INSTANCES">1</field>
            <value name="RESOURCE">
              <block type="aidora_container_resource" id="18">
                <field name="CPU">50</field>
                <field name="MEM">128</field>
              </block>
            </value>
            <value name="IMAGE">
              <block type="aidora_docker_image" id="19">
                <field name="repository">repo</field>
                <field name="name">name</field>
                <field name="tag">latest</field>
              </block>
            </value>
            <next>
              <block type="aidora_container" id="17" inline="false">
                <field name="NAME">default</field>
                <field name="INSTANCES">1</field>
                <value name="RESOURCE">
                  <block type="aidora_container_resource" id="21">
                    <field name="CPU">50</field>
                    <field name="MEM">128</field>
                  </block>
                </value>
                <value name="IMAGE">
                  <block type="aidora_docker_file" id="20">
                    <statement name="NAME">
                      <block type="aidora_docker_from" id="22" inline="true">
                        <value name="NAME">
                          <block type="aidora_docker_image" id="23">
                            <field name="repository">repo</field>
                            <field name="name">name</field>
                            <field name="tag">latest</field>
                          </block>
                        </value>
                        <next>
                          <block type="aidora_docker_net" id="24">
                            <field name="NET">host</field>
                            <next>
                              <block type="aidora_docker_env" id="25">
                                <field name="name">PORT</field>
                                <field name="value">80</field>
                              </block>
                            </next>
                          </block>
                        </next>
                      </block>
                    </statement>
                  </block>
                </value>
              </block>
            </next>
          </block>
        </statement>
        <next>
          <block type="aidora_pod" id="15">
            <statement name="containers">
              <block type="aidora_container" id="26" inline="false">
                <field name="NAME">default</field>
                <field name="INSTANCES">1</field>
              </block>
            </statement>
            <next>
              <block type="aidora_pod" id="27">
                <statement name="containers">
                  <block type="aidora_container" id="28" inline="false">
                    <field name="NAME">default</field>
                    <field name="INSTANCES">1</field>
                  </block>
                </statement>
              </block>
            </next>
          </block>
        </next>
      </block>
    </statement>
  </block>
</xml>`
