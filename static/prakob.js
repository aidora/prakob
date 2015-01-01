/**
    Prakob's Blocks
    (c) 2014-2015 SUT
    Original authors:
    Chanwit Kaewkasi <chanwit@gmail.com>
**/

Blockly.Blocks['aidora_cluster'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(200);
    this.appendDummyInput()
        .appendField("cluster");
    this.appendDummyInput()
        .appendField("N =")
        .appendField(new Blockly.FieldTextInput("1"), "INSTANCES");
    this.appendStatementInput("pods")
        .setCheck("aidora_pod");
  }
};

Blockly.Blocks['aidora_pod'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(200);
    this.appendDummyInput()
        .appendField("pod");
    this.appendStatementInput("containers")
        .setCheck("aidora_container");
    this.setPreviousStatement(true, "aidora_pod");
    this.setNextStatement(true, "aidora_pod");
    this.setTooltip('');
  }
};


Blockly.Blocks['aidora_container'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(160);
    this.appendDummyInput()
        .appendField("container")
        .appendField(new Blockly.FieldTextInput("default"), "NAME");
    this.appendDummyInput()
        .appendField("N =")
        .appendField(new Blockly.FieldTextInput("1"), "INSTANCES");
    this.appendValueInput("RESOURCE")
        .setCheck("aidora_resource")
        .appendField("cpu & mem");
    this.appendValueInput("IMAGE")
        .setCheck("aidora_image")
        .appendField("image");
    this.setPreviousStatement(true, "aidora_container");
    this.setNextStatement(true, "aidora_container");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_container_resource'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(50);
    this.appendDummyInput()
        .appendField("cpu")
        .appendField(new Blockly.FieldDropdown([
            ["50","50"],
            ["100","100"],
            ["150","150"],
            ["200","200"],
            ["250","250"],
            ["300","300"],
            ["350","350"],
            ["400","400"]
        ]), "CPU")
        .appendField("%");
    this.appendDummyInput()
        .appendField("mem")
        .appendField(new Blockly.FieldDropdown([
            ["128", "128"],
            ["256", "256"],
            ["512", "512"],
            ["1024", "1024"],
            ["2048", "2048"],
            ["4096", "4096"],
            ["5120", "5120"]
        ]), "MEM")
        .appendField("MB");
    this.setInputsInline(true);
    this.setOutput(true, "aidora_resource");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_image'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(50);
    this.appendDummyInput()
        .appendField("image");
    this.appendDummyInput()
        .appendField(new Blockly.FieldTextInput("repo"), "repository");
    this.appendDummyInput()
        .appendField("/")
        .appendField(new Blockly.FieldTextInput("name"), "name");
    this.appendDummyInput()
        .appendField(":")
        .appendField(new Blockly.FieldTextInput("latest"), "tag");
    this.setInputsInline(true);
    this.setOutput(true, "aidora_image");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_file'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(50);
    this.appendDummyInput()
        .appendField("image (docker file)");
    this.appendStatementInput("NAME");
    this.setOutput(true, "aidora_image");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_from'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(10);
    this.appendValueInput("NAME")
        .setCheck("aidora_image")
        .appendField("from");
    this.setInputsInline(true);
    this.setPreviousStatement(true, "aidora_docker_stmt");
    this.setNextStatement(true, "aidora_docker_stmt");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_net'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(10);
    this.appendDummyInput()
        .appendField("net")
        .appendField(new Blockly.FieldDropdown([
            ['bridge','bridge'],
            ['host', 'host'],
            ['none','none']
        ]), 'NET');
    this.setInputsInline(true);
    this.setPreviousStatement(true, "aidora_docker_stmt");
    this.setNextStatement(true, "aidora_docker_stmt");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_port'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(10);
    this.appendDummyInput()
        .appendField("port");
    this.appendDummyInput()
        .appendField(new Blockly.FieldTextInput("80"), "host");
    this.appendDummyInput()
        .appendField(":");
    this.appendDummyInput()
        .appendField(new Blockly.FieldTextInput("80"), "guest");
    this.setInputsInline(true);
    this.setPreviousStatement(true, "aidora_docker_stmt");
    this.setNextStatement(true, "aidora_docker_stmt");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_link'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(10);
    this.appendDummyInput()
        .appendField("link");
    this.appendDummyInput()
        .appendField(new Blockly.FieldTextInput("default"), "container");
    this.appendDummyInput()
        .appendField("[");
    this.appendDummyInput()
        .appendField(new Blockly.FieldTextInput("*"), "mul");
    this.appendDummyInput()
        .appendField("]");
    this.setInputsInline(true);
    this.setPreviousStatement(true, "aidora_docker_stmt");
    this.setNextStatement(true, "aidora_docker_stmt");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_volume'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(10);
    this.appendDummyInput()
        .appendField("volume")
        .appendField(new Blockly.FieldTextInput("/"), 'from')
        .appendField(":")
        .appendField(new Blockly.FieldTextInput("/"), 'to');
    this.setInputsInline(true);
    this.setPreviousStatement(true, "aidora_docker_stmt");
    this.setNextStatement(true, "aidora_docker_stmt");
    this.setTooltip('');
  }
};

Blockly.Blocks['aidora_docker_env'] = {
  init: function() {
    // this.setHelpUrl('http://www.example.com/');
    this.setColour(10);
    this.appendDummyInput()
        .appendField("env")
        .appendField(new Blockly.FieldTextInput("ENV"), 'name')
        .appendField("=")
        .appendField(new Blockly.FieldTextInput(""), 'value');
    this.setInputsInline(true);
    this.setPreviousStatement(true, "aidora_docker_stmt");
    this.setNextStatement(true, "aidora_docker_stmt");
    this.setTooltip('');
  }
};