package main

import (
	"fmt"
	"math"
	"regexp"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "测试字符长度\n，谁谁谁sssss"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
	fmt.Println(GetStrLength(str))
}

// GetStrLength 返回输入的字符串的字数，汉字和中文标点算 1 个字数，英文和其他字符 2 个算 1 个字数，不足 1 个算 1个
func GetStrLength(str string) float64 {
	var total float64

	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、/")

	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 1
		} else {
			total = total + 0.5
		}
	}

	return math.Ceil(total)
}

func valid() {
	var a *[]int64

	if len(*a) > 0 {
		fmt.Println(111)
	} else {
		fmt.Println(22)
	}
	fmt.Println()
}
