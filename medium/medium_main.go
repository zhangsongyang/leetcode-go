package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/zhangsongyang/leetcode-go/medium/problem"
)

func main() {
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Equal("he", "he"))

	//142. 环形链表 II
	problem.DetectCycle_main()
	//11. 盛最多水的容器
	problem.MaxArea_main()
}
