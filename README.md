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

    - [`ArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#ArrayList): The implementation of List based on Go built-in slice structure.

    - [`LinkedList`](https://pkg.go.dev/github.com/ghosind/collection/list#LinkedList): The implementation of List based on doubly linked list.

    - [`CopyOnWriteArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#CopyOnWriteArrayList): The thread safe implementation of List based on copy-on-write strategy.

    - [`Stack`](https://pkg.go.dev/github.com/ghosind/collection/list#Stack): The stack implementation based on ArrayList.

- `Set`: A collection interface that contains no duplicate elements.

    - [`HashSet`](https://pkg.go.dev/github.com/ghosind/collection/set#HashSet): The implementation of Set based on Go built-in map structure.

    - [`SyncSet`](https://pkg.go.dev/github.com/ghosind/collection/set#SyncSet): The thread safe implementation of Set based on `sync.Map`.

    - [`DictSet`](https://pkg.go.dev/github.com/ghosind/collection/set#DictSet): The thread safe Set based on RWMutex.

- `Dict`: A object that maps keys to values, and it cannot contain duplicate key.

    - [`HashDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#HashDict): The implementation of Dictionary based on Go built-in map structure.

    - [`SyncDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#SyncDict): The thread safe implementation of dictionary based on `sync.Map`.

    - [`DictDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#DictDict): The thread safe dictionary based on RWMutex.

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
BenchmarkHashDictGet-8          74139873                15.95 ns/op
BenchmarkHashDictPut-8          34336933                31.73 ns/op
BenchmarkLockDictGet-8          14385025                84.57 ns/op
BenchmarkLockDictPut-8          10031228               119.8 ns/op
BenchmarkSyncDictGet-8          191864160                5.795 ns/op
BenchmarkSyncDictPut-8           9078417               129.7 ns/op
```

Set benchmarks with `Add` and `Contains`:

```
BenchmarkHashSet-8      65497208                20.00 ns/op
BenchmarkLockSet-8       9549130               127.4 ns/op
BenchmarkSyncSet-8      61220974                20.90 ns/op
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
