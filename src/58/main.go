package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	arr := strings.Split(strings.TrimSpace(s), " ")
	fmt.Println(arr)
	return len(arr[len(arr) - 1])
}

func main(){
    fmt.Println("Hello, world!");
	result := lengthOfLastWord("   a i l    aaa   ")
	fmt.Println(result);
}
