package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main()  {

}

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:",r.Method)
	if r.Method=="GET" {
		NOW_TIME := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(NOW_TIME, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t,_ := template.ParseFiles("index.html")
		t.Execute(w,token)
	} else {
		r.ParseMultipartForm(32 << 20)  //位运算
		file, handler,err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v",handler.Header)
		f,err:=os.OpenFile()
	}
}
