package main

import "log"

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

}
func main() {
	log.Println("convee's blog:", "http://convee.cn")
	log.Printf("convee:%s\n", "convee.cn")
}
