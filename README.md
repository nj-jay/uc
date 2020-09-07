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

go get -u github.com/nj-jay/uc@v0.1.4

相比v0.0.7　增加的功能为

cal.Play("xx.mp3")可以播放音乐

## example

```
package main

import (

    "fmt"

    "github.com/nj-jay/uc/cal"
)
func main() {

    var a int32
    var b int64
    a = 5
    b = 4
    fmt.Println(cal.Add(a, a))
    fmt.Println(cal.Add(b, b))
    fmt.Println(cal.Add(1, 2, 3))
    fmt.Println(cal.Add(1, 2))
    fmt.Println(cal.Add(1.1, 2.2))
    fmt.Println(cal.Add(2.22, 3.33))
    fmt.Println(cal.Add("I", " ", "love", " ", "go!"))
    fmt.Println(cal.Add("https://", "nj-jay.com"))
    fmt.Println(cal.Add([]int{1, 2}, []int{3, 4}))
 
}

```

运行结果:

10

8

6

3

3.3

5.55

I love go!

https://nj-jay.com

[1 2 3 4]

## 已完成

- [x] Add --int float64 string完成
- [x] Add --[]int完成
- [x] Add --int64 int32完成
- [x] PlayAudio() 实现播放音乐

## Contributors

[nj-jay](https://github.com/nj-jay)
[xmoon5](https://github.com/xmoon5)
