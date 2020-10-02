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
	fmt.Println("9. 回文数:", problem.IsPalindrome(121))
	// 191. 位1的个数
	fmt.Println("191. 位1的个数:", problem.HammingWeight(011))
}
