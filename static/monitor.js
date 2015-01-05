var app = angular.module('monitor', [/*'ui.select'*/]);

app.controller('MonitorCtrl', function ($scope, $http) {

    var url = "http://192.168.1.3:8888/";
    $scope.containers = []

    $scope.updateContainers = function() {

        $http.get(url + "containers/json").
        success(function(data){
            $scope.containers = data;
        });

    };

    setTimeout($scope.updateContainers, 500);
    setInterval($scope.updateContainers, 1500);

    $scope.pause = function(containerId, status) {
        var paused = status.endsWith("(Paused)");
        var cmd = paused ? "unpause" : "pause";

        $http.post(url + "containers/" + containerId + "/" + cmd).
        success(function(data){
            $scope.updateContainers();
        });
    };

    $scope.kill = function(containerId) {
        $http.post(url + "containers/" + containerId + "/kill").
        success(function(data){
            $scope.updateContainers();
        });
    };

})