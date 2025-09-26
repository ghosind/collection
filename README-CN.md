# collection

![Test](https://github.com/ghosind/collection/workflows/collection/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ghosind/collection)](https://goreportcard.com/report/github.com/ghosind/collection)
[![codecov](https://codecov.io/gh/ghosind/collection/branch/main/graph/badge.svg)](https://codecov.io/gh/ghosind/collection)
[![Latest version](https://img.shields.io/github/v/release/ghosind/collection?include_prereleases)](https://github.com/ghosind/collection)
![License Badge](https://img.shields.io/github/license/ghosind/collection)
[![Go Reference](https://pkg.go.dev/badge/github.com/ghosind/collection.svg)](https://pkg.go.dev/github.com/ghosind/collection)

[English](README.md) | 中文

Golang泛型集合框架。

> [!注意]
> 本包需要 Go 1.18 及以上版本。

## 概述

本包提供以下数据结构接口及实现：

- `Collection`：大多数结构的根接口（不包括 `Dict`）。

- `List`：有序集合（也称为序列）。

    - [`ArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#ArrayList)：基于 Go 内置切片结构的 List 实现。

    - [`LinkedList`](https://pkg.go.dev/github.com/ghosind/collection/list#LinkedList)：基于双向链表的 List 实现。

    - [`CopyOnWriteArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#CopyOnWriteArrayList)：基于写时复制策略的线程安全 List 实现。

- `Set`：不包含重复元素的集合接口。

    - [`HashSet`](https://pkg.go.dev/github.com/ghosind/collection/set#HashSet)：基于 Go 内置 map 结构的 Set 实现。

    - [`SyncSet`](https://pkg.go.dev/github.com/ghosind/collection/set#SyncSet)：基于 `sync.Map` 的线程安全 Set 实现。

- `Dict`：将键映射到值的对象，不能包含重复键。

    - [`HashDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#HashDict)：基于 Go 内置 map 结构的字典实现。

    - [`SyncDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#SyncDict)：基于 `sync.Map` 的线程安全字典实现。

## 安装

可以通过以下命令安装本包：

```sh
go get -u github.com/ghosind/collection
```

安装后，可通过如下方式导入：

```go
import "github.com/ghosind/collection"
```

## 示例

### ArrayList 示例

创建一个整数列表，添加并获取元素：

```go
// import "github.com/ghosind/collection/list"

l := list.NewArrayList[int]()
l.Add(10)
l.Add(20)
l.Add(30)
log.Print(l.Get(1)) // 20
```

### HashSet 示例

创建一个字符串集合，添加并判断元素：

```go
// import "github.com/ghosind/collection/set"

fruits := set.NewHashSet[string]()

fruits.Add("Apple")
fruits.Add("Banana")

log.Print(fruits.Contains("Banana")) // true
log.Print(fruits.Contains("Lemon")) // false
```

### HashDict 示例

```go
// import "github.com/ghosind/collection/dict"

languages := dict.NewHashDict[string, int]()

languages.Put("C", 1972)
languages.Put("Go", 2007)

log.Print(languages.GetDefault("C", 0)) // 1972
```

## 许可证

本项目采用 MIT 许可证，详情请参见 [LICENSE](LICENSE) 文件。
