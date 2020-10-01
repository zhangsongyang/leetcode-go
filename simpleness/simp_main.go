package main

import (
	"fmt"
	"github.com/zhangsongyang/leetcode-go/simpleness/problem"

	"github.com/google/go-cmp/cmp"
)

func main() {

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Equal("he", "he"))

	// 9. 回文数
	fmt.Println(problem.IsPalindrome(121))
}
