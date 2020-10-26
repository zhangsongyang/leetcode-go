package problem

import "fmt"

/**
11. 盛最多水的容器
难度
中等

1945

收藏

分享
切换为英文
接收动态
反馈
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

示例：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
*/
func MaxArea(height [9]int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		area := Min(height[left], height[right]) * (right - left)
		ans = Max(ans, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxArea_main() {
	var height = [...]int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println("11. 盛最多水的容器:", MaxArea(height))

}
