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

    - [`list.ArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#ArrayList)：基于 Go 内置切片结构的列表实现。

    - [`list.LinkedList`](https://pkg.go.dev/github.com/ghosind/collection/list#LinkedList)：基于双向链表的列表实现。

    - [`list.CopyOnWriteArrayList`](https://pkg.go.dev/github.com/ghosind/collection/list#CopyOnWriteArrayList)：基于写时复制策略的线程安全列表实现。

    - [`list.LockList`](https://pkg.go.dev/github.com/ghosind/collection/list#LockList)：基于 RWMutex 的线程安全列表包装器。

- `Stack`：遵循后进先出（LIFO）原则的集合。

    - [`stack.Stack`](https://pkg.go.dev/github.com/ghosind/collection/stack#Stack)：基于 ArrayList 的栈实现。

- `Set`：不包含重复元素的集合接口。

    - [`set.HashSet`](https://pkg.go.dev/github.com/ghosind/collection/set#HashSet)：基于 Go 内置 map 结构的集合实现。

    - [`set.SyncSet`](https://pkg.go.dev/github.com/ghosind/collection/set#SyncSet)：基于 `sync.Map` 的线程安全集合实现。

    - [`set.LockSet`](https://pkg.go.dev/github.com/ghosind/collection/set#LockSet)：基于 RWMutex 的线程安全集合包装器。
- `Dict`：将键映射到值的对象，不能包含重复键。

    - [`dict.HashDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#HashDict)：基于 Go 内置 map 结构的字典实现。

    - [`dict.SyncDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#SyncDict)：基于 `sync.Map` 的线程安全字典实现。

    - [`dict.LockDict`](https://pkg.go.dev/github.com/ghosind/collection/dict#LockDict)：基于 RWMutex 的线程安全字典包装器。

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

### 线程安全包装器

这个示例展示了如何使用 `list.LockList` 来创建一个线程安全的列表：

```go
unsafeList := list.NewArrayList[int]()
safeList := list.NewLockList[int](unsafeList)

safeList.Add(10)

log.Print(safeList.Get(0)) // 10
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
BenchmarkBuiltinMap_Get-8               78146028                17.79 ns/op            0 B/op          0 allocs/op
BenchmarkBuiltinMap_Put-8               43320583                25.11 ns/op            0 B/op          0 allocs/op
BenchmarkHashDict_Get-8                 70750198                16.34 ns/op            0 B/op          0 allocs/op
BenchmarkHashDict_Put-8                 34192296                31.89 ns/op            0 B/op          0 allocs/op
BenchmarkLockedHashDict_Get-8           10532293               113.8 ns/op             0 B/op          0 allocs/op
BenchmarkLockedHashDict_Put-8            9389860               129.7 ns/op             0 B/op          0 allocs/op
BenchmarkSyncDict_Get-8                 202971288                6.154 ns/op           0 B/op          0 allocs/op
BenchmarkSyncDict_Put-8                  9171009               130.0 ns/op            16 B/op          1 allocs/op
```

Set 分别执行`Add`/`Contains`基准结果：

```
BenchmarkHashSet_Add-8                  87384132                13.21 ns/op            0 B/op          0 allocs/op
BenchmarkHashSet_Contains-8             87521766                13.77 ns/op            0 B/op          0 allocs/op
BenchmarkLockHashSet_Add-8              41700088                29.54 ns/op            0 B/op          0 allocs/op
BenchmarkLockHashSet_Contains-8         60127644                19.93 ns/op            0 B/op          0 allocs/op
BenchmarkBuiltinMapAsSet_Add-8          87572594                13.49 ns/op            0 B/op          0 allocs/op
BenchmarkBuiltinMapAsSet_Contains-8     84940970                12.69 ns/op            0 B/op          0 allocs/op
BenchmarkSyncSet_Add-8                  10465891               110.3 ns/op             0 B/op          0 allocs/op
BenchmarkSyncSet_Contains-8             245409312                5.671 ns/op           0 B/op          0 allocs/op
```

## 许可证

本项目采用 MIT 许可证，详情请参见 [LICENSE](LICENSE) 文件。
