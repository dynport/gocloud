app = angular.module 'app', ['ngRoute', 'controllers']
app.config ['$routeProvider', '$locationProvider', ($routeProvider, $locationProvider) ->
  $locationProvider.html5Mode(true)
  $routeProvider.
      when("/vpcs/:id",{ templateUrl: "/_vpcs_show.html", controller: 'VpcsShow' }).
      when("/vpcs",{ templateUrl: "/_vpcs_index.html", controller: 'VpcsIndex' }).

      when("/instances/:id",{ templateUrl: "/_instances_show.html", controller: 'InstancesShow' }).
      when("/instances",{ templateUrl: "/_index.html", controller: 'InstancesIndex' }).

      when("/images",{ templateUrl: "/_images_list.html", controller: 'ImagesIndex' }).

      when("/volumes/:id",{ templateUrl: "/_volumes_show.html", controller: 'VolumesShow' }).
      when("/volumes",{ templateUrl: "/_volumes_list.html", controller: 'VolumesIndex' }).

      when("/snapshots/:id",{ templateUrl: "/_snapshots_show.html", controller: 'SnapshotsShow' }).
      when("/snapshots",{ templateUrl: "/_snapshots_list.html", controller: 'SnapshotsIndex' }).

      when("/stacks/:id",{ templateUrl: "/_stacks_show.html", controller: 'StacksShow' }).
      when("/stacks",{ templateUrl: "/_stacks.html", controller: 'StacksIndex' }).

      when("/sgs/:id",{ templateUrl: "/_security_groups.html", controller: 'SecurityGroupsShow' }).
      when("/sgs",{ templateUrl: "/_security_groups_index.html", controller: 'SecurityGroupsIndex' }).
      otherwise({ templateUrl: "/_index.html", controller: 'InstancesIndex' })
]

controllers = angular.module 'controllers', []

controllers.controller 'SnapshotsShow', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/snapshots/" + $routeParams["id"] + ".json").success (snapshot) -> $scope.snapshot = snapshot
]

controllers.controller 'SnapshotsIndex', ["$scope", "$http", ($scope, $http) ->
  $http.get("/api/snapshots.json").success (snapshots) -> $scope.snapshots = snapshots
]

controllers.controller 'VolumesShow', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/volumes/" + $routeParams["id"] + ".json").success (volume) -> $scope.volume = volume
]

controllers.controller 'VolumesIndex', ["$scope", "$http", ($scope, $http) ->
  $http.get("/api/volumes.json").success (volumes) -> $scope.volumes = volumes
]

controllers.controller 'ImagesIndex', ["$scope", "$http", ($scope, $http) ->
  $http.get("/api/images/self.json").success (images) -> $scope.selfImages = images
  $scope.selfImages = []
]

controllers.controller 'VpcsShow', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/vpcs/" + $routeParams["id"] + ".json").success (vpc) -> $scope.vpc = vpc
  $http.get("/api/vpcs/" + $routeParams["id"] + "/subnets.json").success (subnets) -> $scope.subnets = subnets
]

controllers.controller 'VpcsIndex', ["$scope", "$http", ($scope, $http) ->
  $http.get("/api/vpcs.json").success (vpcs) -> $scope.vpcs = vpcs
]

controllers.controller 'StacksIndex', ["$scope", "$http", ($scope, $http) ->
  $http.get("/api/stacks.json").success (stacks) -> $scope.stacks = stacks
]

controllers.controller 'StacksShow', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/stacks/" + $routeParams["id"] + ".json").success (stack) ->
    $scope.stack = stack
  $http.get("/api/stacks/" + $routeParams["id"] + "/resources.json").success (res) ->
    $scope.resources = res
]

controllers.controller 'SecurityGroupsIndex', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/security_groups.json").success (groups) -> $scope.groups = groups
]

controllers.controller 'SecurityGroupsShow', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/security_groups/" + $routeParams.id + ".json").success (group) -> $scope.group = group
  $http.get("/api/security_groups/" + $routeParams.id + "/instances.json").success (instances) -> $scope.instances = instances
]

controllers.controller 'InstancesIndex', ["$scope", "$http", ($scope, $http) ->
  $http.get("/api/instances.json").success (instances) -> $scope.instances = instances
]

controllers.controller 'InstancesShow', ["$scope", "$http", "$routeParams", ($scope, $http, $routeParams) ->
  $http.get("/api/instances/" + $routeParams.id + ".json").success (instance) -> $scope.instance = instance
]
