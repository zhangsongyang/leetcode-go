package main

import (
	"fmt"
	"gitee.com/go-package/carbon"
	"github.com/google/go-cmp/cmp"
	"github.com/zhangsongyang/leetcode-go/simpleness/problem"
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
	fmt.Println("输入一个int:")
	var n int
	fmt.Scanln(&n)
	fmt.Printf("231. 2的幂--> %v是否是 2 的幂次:%v\n", n, problem.IsPowerOfTwo(n))
	// 190. 颠倒二进制位
	fmt.Printf("190. 颠倒二进制位:%b", problem.ReverseBits(4294967293))

}
