package algs

/**
 * Definition for a binary tree node.
 * Binary tree
 * 二叉树节点
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for singly-linked list.
 * 单链表节点 
 */
type ListNode struct {
	Val int
	Next *ListNode
}