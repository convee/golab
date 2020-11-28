package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

func main() {

	// 计算密码MD5
	c := md5.New()
	io.WriteString(c, "BBA2CABF-1899-40A0-824E-00E8D1E70B90")
	spw := fmt.Sprintf("%x", c.Sum(nil))

	// 指定两个(salt)
	salt1 := "rA32iv{oAu6WHxtXu@MP"
	salt2 := "47ebyEXn2TpwaYvFF;*R"

	// 拼接密码MD5
	buf := bytes.NewBufferString("")

	// 拼接密码
	io.WriteString(buf, salt1)
	io.WriteString(buf, spw)
	io.WriteString(buf, salt2)

	// 拼接密码计算MD5
	t := md5.New()
	io.WriteString(t, buf.String())

	fmt.Printf("salt1 : %s , spw : %s , salt2 : %s \n", salt1, spw, salt2)

	// 输出
	fmt.Printf("%x\n", t.Sum(nil))
}
