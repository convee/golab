package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fileName := "target.tar.gz"
	var r *bufio.Reader
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, can not open %s: error:%s\n", os.Args[0], fileName, err)
		os.Exit(1)
	}
	fz, err := gzip.NewReader(fi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open gzip failed,err:%v\n", err)
		return
	}
	r = bufio.NewReader(fz)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
