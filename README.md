# Using-Cal-go

## function

仅用一个函数add实现的功能

理想的结果如下

```
//支持int int32 int64 string float32 float64类型
add(1, 2) = 3
add(1.1, 2.2) = 3.3
add("hello", "world") = "hello world"
add([]int{1, 2}, []int{4, 5, 6}) = []int{1, 2, 4, 5, 6}
add([]float32{1.1, 2.2}, []float32{3.3, 4.4}) = []float32{1.1, 2.2, 3.3, 4.3}

//后续需要用到哪些再进行补充
```

## How-to-Use

go get -u github.com/nj-jay/uc@v0.0.2

## example

```go
package main

import (

	"fmt"

	"github.com/nj-jay/uc"

)

func main(){

    fmt.Println(uc.Add(1, 2))

	fmt.Println(uc.Add(1.1, 2.2))

    fmt.Println(uc.Add("I", " ", "love", " ", "go!"))

}

```

运行结果:

3

3.3

1 love go!
## 已完成

- [x] Add --int float64 string完成
- [x] Add --[]int完成



