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
	//144. 二叉树的前序遍历
	problem.PreorderTraversal_main()
	//198. 打家劫舍
	problem.Rob_main()
	//64. 最小路径和
	problem.MinPathSum_main()
	//238. 除自身以外数组的乘积
	problem.ProductExceptSelf_main()
}
