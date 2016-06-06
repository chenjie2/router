var app = angular.module('myApp', []);
app.controller('personCtrl', function($scope, $http) {
	$http.get("http://127.0.0.1:8888/api?command=GetDomains").success(
		function(response) {
			$scope.names = response;
		}
	);
	$scope.domain='test111';
	$scope.upstream = 'server 127.0.0.1;'
	$scope.deleteDomain = "";
	$scope.publish = function() {
		var domain = $scope.domain;
		var server = $scope.upstream;
		$http.post('http://127.0.0.1:8888/api', {
                command: 'AddDomain',
                domain: domain,
                server: server
            }).success(function(response){
            	var newitem = {DomainName:domain,Server:server}
            	$scope.names.push(newitem)
            });
    };
	$scope.delete = function(domainName) {
        $scope.deleteDomain = domainName;
        $('#deletedomain').modal('toggle');
    };
    $scope.deleteAction = function() {
    	var domain = $scope.deleteDomain
		$http.post('http://127.0.0.1:8888/api', {
                command: 'DelDomain',
                domain: domain
            }).success(function(response){
            	$('#'+domain).remove();
            });
		$('#deletedomain').modal('toggle');
    };
});