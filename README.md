# collection

![Test](https://github.com/ghosind/collection/workflows/collection/badge.svg)
[![codecov](https://codecov.io/gh/ghosind/collection/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/collection)
[![Latest version](https://img.shields.io/github/v/release/ghosind/collection?include_prereleases)](https://github.com/ghosind/collection)
![License Badge](https://img.shields.io/github/license/ghosind/collection)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/collection.svg)](https://pkg.go.dev/github.com/ghosind/collection)

Generics collections framework for Golang.

> [!NOTE]
> This package requires Go version 1.18+.

## Overview

This package provides the following data structure interfaces and implementations:

- `Collection`: The root interface of most of the structures in this package (without `Dictionary`).

- `List`: An ordered collection (also known as a sequence).

    - [`ArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#ArrayList): The implementation of List based on Go built-in slice structure.

- `Set`: A collection interface that contains no duplicate elements.

    - [`HashSet`](https://pkg.go.dev/github.com/ghosind/collection/set#HashSet): The implementation of Set based on Go built-in map structure.

    - [`SyncSet`](https://pkg.go.dev/github.com/ghosind/collection/set#SyncSet): The thread safe implementation of Set based on `sync.Map`.

- `Dict`: A object that maps keys to values, and it cannot contain duplicate key.

    - [`HashDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#HashDict): The implementation of Dictionary based on Go built-in map structure.

    - [`SyncDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#SyncDict): The thread safe implementation of dictionary based on `sync.Map`.

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

### HashSet Examples

Create a string set, add and test elements in the set.

```go
// import "github.com/ghosind/collection/set"

fruits := set.NewHashSet[string]()

fruits.Add("Apple")
fruits.Add("Banana")

log.Print(fruits.Contains("Banana")) // true
log.Print(fruits.Contains("Lemon"))
```

### HashDict Examples

```go
// import "github.com/ghosind/collection/dict"

languages := dict.NewHashDict[string, int]()

languages.Put("C", 1972)
languages.Put("Go", 2007)

log.Print(languages.GetDefault("C", 0)) // 1972
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
