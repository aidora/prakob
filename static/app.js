var app = angular.module('app', [/*'ui.select'*/]);

app.controller('AppCtrl', function ($scope, $http) {

  $scope.refresh = function() {
    $http.get("/filenames").
    success(function(data, status, header, config) {
      $scope.filenames = data['filenames'];
    });
  }

  $scope.refresh()

  $scope.doSave = function() {
    var wk = window.prompt("Save as", $scope.workspace);
    var dom = Blockly.Xml.workspaceToDom(Blockly.mainWorkspace);
    var xmlText = Blockly.Xml.domToPrettyText(dom);

    $http.post("/save",
       "xml=" + xmlText + "&" + "workspace=" + wk,
       {headers: {'Content-Type': 'application/x-www-form-urlencoded'}}).
    success(function(data, status, header, config) {
      $scope.refresh()
    });
  };

  $scope.doLoad = function() {
    $http.get("/load/" + $scope.workspace).
    success(function(data, status, header, config) {
      Blockly.mainWorkspace.clear();
      var xml = Blockly.Xml.textToDom(data);
      Blockly.Xml.domToWorkspace(Blockly.mainWorkspace, xml);
    });
  };

  $scope.doClear = function() {
   Blockly.mainWorkspace.clear();
  };

  $scope.doDeploy = function() {
    $http.post("/deploy/" + $scope.workspace).
    success(function(data, status, header, config) {
      alert(data);
    });
  };

});