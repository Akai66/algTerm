package main

import (
	"algTerm/string/term"
	"fmt"
)

//测试字符串相关算法题

func main() {
	//测试Str2int
	fmt.Println(term.Str2int(""))
	fmt.Println(term.Str2int("-"))
	fmt.Println(term.Str2int("123a4"))
	fmt.Println(term.Str2int("+123"))
	fmt.Println(term.Str2int("-456"))
	fmt.Println(term.Str2int("2147483647"))
	fmt.Println(term.Str2int("2147483648"))

	//============================================

}
