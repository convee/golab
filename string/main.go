package main

import "fmt"

func main() {
	traversalString()
	changeString()
}

// 遍历字符串
// UTF8编码下一个中文汉字由3~4个字节组成，
// 所以我们不能简单的按照字节去遍历一个包含中文的字符串，
// 否则就会出现上面输出中第一行的结果
func traversalString() {
	s := "convee.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

// 修改字符串
// 要修改字符串，需要先将其转换成[]rune或[]byte，
// 完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组
func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}
