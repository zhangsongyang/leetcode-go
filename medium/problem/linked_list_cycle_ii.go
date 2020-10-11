package problem

import "fmt"

/**
142. 环形链表 II
难度
中等

691

收藏

分享
切换为英文
接收动态
反馈
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。



示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：

输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。


示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。


进阶：
你是否可以不用额外空间解决此题？

https://leetcode-cn.com/problems/linked-list-cycle-ii/
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func DetectCycle_main() {
	var head ListNode
	var headB ListNode
	var headC ListNode
	var headD ListNode
	head.Val = 3
	headB.Val = 2
	headC.Val = 0
	headD.Val = -4
	head.Next = &headB
	headB.Next = &headC
	headC.Next = &headD
	headD.Next = &headB

	fmt.Println(DetectCycle(&head))
}
