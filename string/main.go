package main

import (
	"algTerm/string/term"
	"fmt"
)

//test string alg

func main() {
	fmt.Println("============================================")
	fmt.Println("test Str2int")
	fmt.Println("============================================")
	fmt.Println(term.Str2int(""))
	fmt.Println(term.Str2int("-"))
	fmt.Println(term.Str2int("123a4"))
	fmt.Println(term.Str2int("+123"))
	fmt.Println(term.Str2int("-456"))
	fmt.Println(term.Str2int("2147483647"))
	fmt.Println(term.Str2int("2147483648"))

	fmt.Println("============================================")
	fmt.Println("test kmpSearch")
	fmt.Println("============================================")

	fmt.Println(term.KmpSearch("adfeababcbcab", "cbcab"))
}
