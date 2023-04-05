package binarytree

import (
	"fmt"
	"time"
)

// 中序 递归
func Inorder(root *ListNode) []int {
	var res []int
	order := func(*ListNode) {}

	order = func(r *ListNode) {
		if r == nil {
			return
		}
		order(r.Left)
		res = append(res, r.Data)
		order(r.Right)
	}

	order(root)

	return res
}

// 中序 非递归
func Inorder2(root *ListNode) []int {
	var res []int
	var stack []*ListNode
	node := root

	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Data)
		node = node.Right
	}

	return res
}

// 后续 递归
func PostOrder(root *ListNode) []int {
	var res []int
	order := func(r *ListNode) {}
	order = func(r *ListNode) {
		if r == nil {
			return
		}
		order(r.Left)
		order(r.Right)
		res = append(res, r.Data)
	}
	order(root)
	return res
}

// 后序 非递归
func PostOrder2(root *ListNode) []int {
	var res []int
	var stack []*ListNode
	node := root
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || node == root.Right {
			res = append(res, root.Data)
			node = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
		fmt.Println(stack)
		fmt.Println(res)
		time.Sleep(1)
	}

	return res
}
