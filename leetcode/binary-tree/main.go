package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	list := []int{}
	traversalBefore(root, &list)
	return list
}

//前序遍历
func traversalBefore(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	*list = append(*list, root.Val)
	traversalBefore(root.Left, list)
	traversalBefore(root.Right, list)
}

//中序遍历
func traversalMiddle(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	traversalMiddle(root.Left, list)
	*list = append(*list, root.Val)
	traversalMiddle(root.Right, list)
}

//后序遍历
func traversalAfter(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	traversalAfter(root.Left, list)
	traversalAfter(root.Right, list)
	*list = append(*list, root.Val)
}

func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)

	q := make(chan TreeNode, 100)

	if root != nil {
		q <- *root
	}

	for len(q) != 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := <-q
			if cur.Left != nil {
				q <- *(cur.Left)
			}
			if cur.Right != nil {
				q <- *(cur.Right)
			}

			if i == size-1 {
				ret = append(ret, cur.Val)
			}
		}
	}
	return ret
}
func main() {
	root := new(TreeNode)
	root.Val = 1
	right := new(TreeNode)
	right.Val = 2
	rightLeft := new(TreeNode)
	rightLeft.Val = 3
	right.Left = rightLeft
	root.Right = right
	fmt.Println(preorderTraversal(root))
}
