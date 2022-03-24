
该库参考了laravel的集合，PHP转go的可以快速上手。

- [集合](#集合)
  - [安装和使用](#安装和使用)

## 安装和使用
<hr>

1. 使用命令安装集合。

```sh
$ go get -u github.com/greece-olympus/go-collection
```
2. 在你的项目引入。

```go
import collection "github.com/greece-olympus/go-collection"
```

3. 开始使用

```go
package main

import (
  "fmt"
  collection "github.com/greece-olympus/go-collection"
)

func main() {
  col := collection.NewCollection([]int{1, 2, 3, 4})
  fmt.Println(col.Count())
}
```
4. 运行后输出
```
4

进程 已完成，退出代码为 0
```

## 安装和使用
<hr>