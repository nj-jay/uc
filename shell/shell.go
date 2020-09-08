package shell

import (

    "fmt"
    "net/http"
    "os"
    "io"
    "strings"
    "log"
    "strconv"

)

func Wget(url string) {
  
  	//打印时间
  	log.Println(url)
  	
  	res, _ := http.Get(url)
  	
  	fmt.Print("已发出" + res.Proto + "请求, 正在等待回应...")
  	
  	//判断Status
  	if res.StatusCode == 404 {
  	
  		fmt.Println(res.Status)
  		
  		os.Exit(0)
  		
  	} else {
  	
  		fmt.Println(res.Status)
  		
		index := strings.LastIndex(url, "/")
		
		
		//当前目录创建文件
    	f, _ := os.Create(url[index+1:])
    	
    	length := res.Header.Get("Content-Length")
    	
   	 	length_kb, _ := strconv.ParseFloat(length, 64)
   	 	
   		fmt.Println("长度:	", "[", length, "byte", "]", ">--", fmt.Sprintf("%.2f", length_kb / 1024), "Kb", "	[", res.Header.Get("Content-Type"), "]")
   		
   		fmt.Println("正在保存至:", url[index+1:])
   		
   		//写到f中
    	io.Copy(f, res.Body)
    	
    	log.Println("已保存", url[index+1:])
    	
    	//关闭连接
    	_ = res.Close
  		
  	}
}

