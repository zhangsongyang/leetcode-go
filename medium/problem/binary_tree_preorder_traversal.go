package problem

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 144. 二叉树的前序遍历
 * 难度
 * 中等
 * <p>
 * 420
 * <p>
 * 收藏
 * <p>
 * 分享
 * 切换为英文
 * 接收动态
 * 反馈
 * 给定一个二叉树，返回它的 前序 遍历。
 * <p>
 * 示例:
 * <p>
 * 输入: [1,null,2,3]
 * 1
 * \
 * 2
 * /
 * 3
 * <p>
 * 输出: [1,2,3]
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 */
func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

func PreorderTraversal_main() {
	var leftTwo *TreeNode = &TreeNode{3, nil, nil}
	var rightOne *TreeNode = &TreeNode{2, leftTwo, nil}
	var root *TreeNode = &TreeNode{1, nil, rightOne}
	fmt.Println("144. 二叉树的前序遍历:", preorderTraversal(root))

}
