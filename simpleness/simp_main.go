package main

import (
	"fmt"
	"github.com/zhangsongyang/leetcode-go/simpleness/problem"

	"gitee.com/go-package/carbon"
	"github.com/google/go-cmp/cmp"
)

func main() {

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Equal("he", "he"))
	c := carbon.New().Timezone(carbon.PRC)
	fmt.Println(c.Today())

	// 9. 回文数
	fmt.Println("9. 回文数:", problem.IsPalindrome(121))
	// 191. 位1的个数
	fmt.Println("191. 位1的个数:", problem.HammingWeight(011))
	// 231. 2的幂
	fmt.Println("231. 2的幂:", problem.IsPowerOfTwo(16))

}
