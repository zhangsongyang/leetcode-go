package main

import (
	"fmt"
	"gitee.com/go-package/carbon"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/zhangsongyang/leetcode-go/simpleness/problem"
)

func main() {

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Equal("he", "he"))
	c := carbon.New().Timezone(carbon.PRC)
	fmt.Println(c.Today())
	fmt.Println(uuid.New().String())

	// 9. 回文数
	fmt.Println("9. 回文数:", problem.IsPalindrome(121))
	// 191. 位1的个数
	fmt.Println("191. 位1的个数:", problem.HammingWeight(011))
	// 231. 2的幂
	fmt.Println("输入一个int:")
	//var n int
	//fmt.Scanln(&n)
	n := 25
	fmt.Printf("231. 2的幂--> %v是否是 2 的幂次:%v\n", n, problem.IsPowerOfTwo(n))
	// 190. 颠倒二进制位
	fmt.Printf("190. 颠倒二进制位:%b\n", problem.ReverseBits(4294967293))
	// 925.长按键入
	fmt.Printf("925.长按键入: %v\n", problem.IsLongPressedName("alex", "aaleex"))

}
