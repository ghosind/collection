# collection

![Test](https://github.com/ghosind/collection/workflows/collection/badge.svg)
[![codecov](https://codecov.io/gh/ghosind/collection/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/collection)
[![Latest version](https://img.shields.io/github/v/release/ghosind/collection?include_prereleases)](https://github.com/ghosind/collection)
![License Badge](https://img.shields.io/github/license/ghosind/collection)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/collection.svg)](https://pkg.go.dev/github.com/ghosind/collection)

Generics collections framework for Golang.

> IMPORTANT NOTICE: This package requires Go version 1.18+.

## Overview

This package provides the following data structure interfaces and implementations:

- `Collection`: The root interface of most of the structures in this package.

- `Set`: A collection interface that contains no duplicate elements.

    - [`HashSet`](https://pkg.go.dev/github.com/ghosind/collection#HashSet): The implementation of Set based on Go built-in map structure.

    - [`ConcurrentHashSet`](https://pkg.go.dev/github.com/ghosind/collection#ConcurrentHashSet): The thread safe implementation of Set based on Go built-in map structure.

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
fruits := collection.NewHashSet[string]()

fruits.Add("Apple")
fruits.Add("Banana")

log.Print(fruits.Contains("Banana")) // true
log.Print(fruits.Contains("Lemon"))
```
