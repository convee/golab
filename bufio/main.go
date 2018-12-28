package main

import (
	"bufio"
	"fmt"
	"os"
)

func init()  {

}

//计算出现的数量
func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
func dup1()  {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	for line,n:=range counts {
		if n > 1 {
			fmt.Println(fmt.Printf("%d\t%s\n",n,line))
		}
	}
}
func dup2()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _,arg := range files {
			f,err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line,n:=range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}
func main()  {
	//dup1()
	dup2()
}
