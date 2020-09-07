# Using-Concise-Go

## How-to-Use

go get -u github.com/nj-jay/uc@v1.0.0

## Function

目录结构:

* cal库 计算库 

  > 目前已实现的功能
  >
  > import "github.com/nj-jay/uc/cal"
  >
  > * cal.Add(a ...interface{}) 
  >
  >   空接口类型可以为int, int32, int64, string, []int, float64
  >
  >   ### example
  >
  > ```go
  > package main
  > 
  > import (
  > 
  >     "fmt"
  >     "github.com/nj-jay/uc/cal"
  > )
  > func main() {
  > 
  >     var a int32
  >     var b int64
  >     a = 5
  >     b = 4
  >     fmt.Println(cal.Add(a, a))
  >     fmt.Println(cal.Add(b, b))
  >     fmt.Println(cal.Add(1, 2, 3))
  >     fmt.Println(cal.Add(1, 2))
  >     fmt.Println(cal.Add(1.1, 2.2))
  >     fmt.Println(cal.Add(2.22, 3.33))
  >     fmt.Println(cal.Add("I", " ", "love", " ", "go!"))
  >     fmt.Println(cal.Add("https://", "nj-jay.com"))
  >     fmt.Println(cal.Add([]int{1, 2}, []int{3, 4}))
  > }
  > //运行结果
  > 10
  > 8
  > 6
  > 3
  > 3.3
  > 5.55
  > I love go!
  > https://nj-jay.com
  > [1 2 3 4]
  > ```
  >

* play库

  > 目前已实现的功能
  >
  > import "github.com/nj-jay/uc/play"
  >
  > * play.PlayAudio(s string) 
  >
  >   参数为mp3文件的路径
  >
  >   ### example
  >
  >   ```go
  >   package main
  >   
  >   import (
  >   
  >       "fmt"
  >       "github.com/nj-jay/uc/play"
  >   )
  >   
  >   func main() {
  >       play.PlayAudio("./七里香.mp3")
  >   }
  >   //运行结果
  >   播放音乐，音乐播放完退出程序
  >   ```

## 已完成

* cal.Add()
* play.PlayAudio()

## Contributors

[nj-jay](https://github.com/nj-jay)

[xmoon5](https://github.com/xmoon5)
