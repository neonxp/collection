# Collection with generics

Go >= 1.18 required.

# Installation

`go get go.neonxp.dev/collection@latest`

# Methods

|Method|Description|Example|
|:-----|:----------|------:|
|`Map`|Async map over slice|[example_map_test.go](./example_map_test.go)|
|`MapSync`|Sync map over slice|[example_map_test.go](./example_map_test.go)|
|`Each`|Async call cb over each element|[example_each_test.go](./example_each_test.go)|
|`MapEach`|Sync call cb over each element|[example_each_test.go](./example_each_test.go)|
|`Filter`|Returns filtered elements async|TODO|
|`FilterSync`|Returns filtered elements|TODO|
|`Reduce`|Produce one single result from a sequence of elements|TODO|
