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

    - [`list.ArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#ArrayList)：基于 Go 内置切片结构的 List 实现。

    - [`list.LinkedList`](https://pkg.go.dev/github.com/ghosind/collection/list#LinkedList)：基于双向链表的 List 实现。

    - [`list.CopyOnWriteArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#CopyOnWriteArrayList)：基于写时复制策略的线程安全 List 实现。

- `Stack`：遵循后进先出（LIFO）原则的集合。

    - [`list.Stack`](https://pkg.go.dev/github.com/ghosind/collection/list#Stack)：基于 ArrayList 的栈实现。

- `Set`：不包含重复元素的集合接口。

    - [`set.HashSet`](https://pkg.go.dev/github.com/ghosind/collection/set#HashSet)：基于 Go 内置 map 结构的 Set 实现。

    - [`set.SyncSet`](https://pkg.go.dev/github.com/ghosind/collection/set#SyncSet)：基于 `sync.Map` 的线程安全 Set 实现。

    - [`set.DictSet`](https://pkg.go.dev/github.com/ghosind/collection/set#DictSet)：基于 RWMutex 的线程安全 Set 实现。

- `Dict`：将键映射到值的对象，不能包含重复键。

    - [`dict.HashDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#HashDict)：基于 Go 内置 map 结构的字典实现。

    - [`dict.SyncDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#SyncDict)：基于 `sync.Map` 的线程安全字典实现。

    - [`dict.LockDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#LockDict)：基于 RWMutex 的线程安全字典实现。

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

## 测试

运行整个仓库的单元测试：

```sh
go test ./...
```

运行基准测试（所有包）：

```sh
go test -bench=. -benchmem ./...
```

仅运行单个包的基准（例如：`dict`）：

```sh
go test ./dict -bench=. -run=^$ -benchmem
```

## 基准测试（Apple M2 示例结果）

以下为在 Apple M2 机器上得到的示例基准结果，实际结果可能因 Go 版本和系统负载有所不同。

Dict 分别执行`Get`/`Put`基准结果：

```
BenchmarkHashDictGet-8          75120351                16.42 ns/op            0 B/op          0 allocs/op
BenchmarkHashDictPut-8          36379850                36.31 ns/op            0 B/op          0 allocs/op
BenchmarkLockDictGet-8          14418043                82.43 ns/op            0 B/op          0 allocs/op
BenchmarkLockDictPut-8           9442551               120.4 ns/op             0 B/op          0 allocs/op
BenchmarkSyncDictGet-8          202543980                7.756 ns/op           0 B/op          0 allocs/op
BenchmarkSyncDictPut-8           8971009               132.8 ns/op            16 B/op          1 allocs/op
```

Set 分别执行`Add`与`Contains`基准结果：

```
BenchmarkHashSet-8      60629182                22.92 ns/op            0 B/op          0 allocs/op
BenchmarkLockSet-8       9375787               126.8 ns/op             0 B/op          0 allocs/op
BenchmarkSyncSet-8      50990958                23.69 ns/op            0 B/op          0 allocs/op
```

## 许可证

本项目采用 MIT 许可证，详情请参见 [LICENSE](LICENSE) 文件。
