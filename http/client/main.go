package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch()  {
	res, err := http.Get("https://www.baidu.com/")
	if err != nil {
		fmt.Println("get err:", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("get data err:", err)
		return
	}
	fmt.Println(string(data))
}

func fetchOne(url string, ch chan<- string)  {
	start := time.Now()
	resp,err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()
	if err!= nil {
		ch <- fmt.Sprintf("while reading %s:%v",url,err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes,url)
}
func fetchAll()  {
	start := time.Now()
	ch := make(chan string)
	for _,url := range os.Args[1:] {
		go fetchOne(url, ch)//启动一个goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)//从通道ch接收
	}
	fmt.Printf("%.2fs elapsed\n",time.Since(start).Seconds())
}
func main() {
	fetchAll()
}
