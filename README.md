# Using-Concise-Go

## How-to-Use


go get -u github.com/nj-jay/uc@v2.0.0

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
  >
  >       "github.com/nj-jay/uc/play"
  >   )
  >   
  >   func main() {
  >       play.PlayAudio("./七里香.mp3")
  >   }
  >   //运行结果
  >   播放音乐，音乐播放完退出程序
  >   ```

* shell库

    > shell.Wget(url string) 可以下载网络上的文件 致敬wget命令
    >
    > ```go
    > package main
    > import (
    > 	"github.com/nj-jay/uc/shell"
    > )
    > 
    > func main() {
    > shell.Wget("https://pucture.nj-jay.com/node.tar.xz")
    > }
    > //运行结果
    > 2020/10/15 23:37:39 https://picture.nj-jay.com/node.tar.xz
    > 已发出HTTP/2.0请求, 正在等待回应...200 OK
    > 长度:	 [ 14781396 byte ] >-- 14434.96 Kb 	[ application/x-xz ]
    > 正在保存至: node.tar.xz
    >  14.78 MB / 14.78 MB [====================================] 100.00% 5.81 MB/s 2s
    > 2020/10/15 23:37:43 已保存 node.tar.xz
    > ```
    >
    > shell.Ls(path string) 致敬ls命令
    >
    > ```go
    > package main
    > import (
    > 	"github.com/nj-jay/uc/shell"
    > )
    > 
    > func main() {
    >  if len(os.Args) == 1 {
    > 		shell.Ls(".")
    > 	} else {
    > 		path := os.Args[1]
    > 		shell.Ls(path)
    > 	}
    > }
    > // 编译成二进制文件ls-go
    > //运行结果 
    > LICENSE   0.1 Kb
    > README.md   2.5 Kb
    > cal
    > example
    > go.mod   0.1 Kb
    > go.sum   1.1 Kb
    > play
    > shell
    > // ls-go ../
    > Somusic
    > book
    > calcugo
    > cgo
    > faker
    > file
    > fileUpload
    > fyne
    > gif
    > ```



## 已完成

* cal.Add()
* play.PlayAudio()
* shell.Wget()
* shell.Ls()

## 正在开发的功能



## Contributors

[nj-jay](https://github.com/nj-jay)

[xmoon5](https://github.com/xmoon5)
