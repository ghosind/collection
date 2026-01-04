# collection

![Test](https://github.com/ghosind/collection/workflows/collection/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/collection)](https://goreportcard.com/report/github.com/ghosind/collection)
[![codecov](https://codecov.io/gh/ghosind/collection/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/collection)
[![Latest version](https://img.shields.io/github/v/release/ghosind/collection?include_prereleases)](https://github.com/ghosind/collection)
![License Badge](https://img.shields.io/github/license/ghosind/collection)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/collection.svg)](https://pkg.go.dev/github.com/ghosind/collection)

English | [中文](README-CN.md)

Generics collections framework for Golang.

> [!NOTE]
> This package requires Go version 1.18+.

## Overview

This package provides the following data structure interfaces and implementations:

- `Collection`: The root interface of most of the structures in this package (without `Dict`).

- `List`: An ordered collection (also known as a sequence).

    - [`list.ArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#ArrayList): The implementation of List based on Go built-in slice structure.

    - [`list.LinkedList`](https://pkg.go.dev/github.com/ghosind/collection/list#LinkedList): The implementation of List based on doubly linked list.

    - [`list.CopyOnWriteArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#CopyOnWriteArrayList): The thread safe implementation of List based on copy-on-write strategy.

- `Stack`: A collection that follows the LIFO (last-in, first-out) principle.

    - [`list.Stack`](https://pkg.go.dev/github.com/ghosind/collection/list#Stack): The stack implementation based on ArrayList.

- `Set`: A collection interface that contains no duplicate elements.

    - [`set.HashSet`](https://pkg.go.dev/github.com/ghosind/collection/set#HashSet): The implementation of Set based on Go built-in map structure.

    - [`set.SyncSet`](https://pkg.go.dev/github.com/ghosind/collection/set#SyncSet): The thread safe implementation of Set based on `sync.Map`.

    - [`set.DictSet`](https://pkg.go.dev/github.com/ghosind/collection/set#DictSet): The thread safe Set based on RWMutex.

- `Dict`: A object that maps keys to values, and it cannot contain duplicate key.

    - [`dict.HashDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#HashDict): The implementation of Dictionary based on Go built-in map structure.

    - [`dict.SyncDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#SyncDict): The thread safe implementation of dictionary based on `sync.Map`.

    - [`dict.LockDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#LockDict): The thread safe dictionary based on RWMutex.

## Installation

You can install this package by the following command.

```sh
go get -u github.com/ghosind/collection
```

After installation, you can import it by the following code.

```go
import "github.com/ghosind/collection"
```

## Examples

### ArrayList Examples

Create an integer list, add and get elements from the list.

```go
// import "github.com/ghosind/collection/list"

l := list.NewArrayList[int]()
l.Add(10)
l.Add(20)
l.Add(30)
log.Print(l.Get(1)) // 20
```

### HashSet Examples

Create a string set, add and test elements in the set.

```go
// import "github.com/ghosind/collection/set"

fruits := set.NewHashSet[string]()

fruits.Add("Apple")
fruits.Add("Banana")

log.Print(fruits.Contains("Banana")) // true
log.Print(fruits.Contains("Lemon")) // false
```

### HashDict Examples

```go
// import "github.com/ghosind/collection/dict"

languages := dict.NewHashDict[string, int]()

languages.Put("C", 1972)
languages.Put("Go", 2007)

log.Print(languages.GetDefault("C", 0)) // 1972
```

## Testing

Run unit tests for the whole repository:

```sh
go test ./...
```

Run benchmarks (all packages):

```sh
go test -bench=. -benchmem ./...
```

Run benchmarks for a single package (example: `dict`):

```sh
go test ./dict -bench=. -run=^$ -benchmem
```

## Benchmarks (Apple M2 sample results)

Below are sample benchmark results run on an Apple M2 machine. Your results may vary depending on Go version and system load.

Dict benchmarks with `Get`/`Put`:

```
BenchmarkBuiltinMap_Get-8       78598278                15.34 ns/op            0 B/op          0 allocs/op
BenchmarkBuiltinMap_Put-8       48577176                26.72 ns/op            0 B/op          0 allocs/op
BenchmarkHashDict_Get-8         72823633                16.32 ns/op            0 B/op          0 allocs/op
BenchmarkHashDict_Put-8         32907440                34.03 ns/op            0 B/op          0 allocs/op
BenchmarkLockDict_Get-8         14978312                82.18 ns/op            0 B/op          0 allocs/op
BenchmarkLockDict_Put-8          9767184               123.2 ns/op             0 B/op          0 allocs/op
BenchmarkSyncDict_Get-8         192600193                5.701 ns/op           0 B/op          0 allocs/op
BenchmarkSyncDict_Put-8          9013268               129.3 ns/op            16 B/op          1 allocs/op
```

Set benchmarks with `Add`/`Contains`:

```
BenchmarkHashSet_Add-8                  83323072                14.53 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet_Contains-8             90324601                14.28 ns/op            0 B/op          0 allocs/op
BenchmarkLockSet_Add-8                  10818360               113.4 ns/op             0 B/op          0 allocs/op
BenchmarkLockSet_Contains-8             14657197                82.17 ns/op            0 B/op          0 allocs/op
BenchmarkBuiltinMapAsSet_Add-8          92263930                14.61 ns/op            0 B/op          0 allocs/op
BenchmarkBuiltinMapAsSet_Contains-8     98068932                11.20 ns/op            0 B/op          0 allocs/op
BenchmarkSyncSet_Add-8                  10983445               108.0 ns/op             0 B/op          0 allocs/op
BenchmarkSyncSet_Contains-8             272523782                4.427 ns/op           0 B/op          0 allocs/op
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
