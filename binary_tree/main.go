package main

import (
	"fmt"
)

type Tree struct {
	Name  string
	Left  *Tree
	Right *Tree
}

//前序遍历
func trans_before(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans_before(root.Left)
	trans_before(root.Right)
}

//中序遍历
func trans_middle(root *Tree) {
	if root == nil {
		return
	}
	trans_middle(root.Left)
	fmt.Println(root)
	trans_middle(root.Right)
}

//后序遍历
func trans_after(root *Tree) {
	if root == nil {
		return
	}
	trans_after(root.Left)
	trans_after(root.Right)
	fmt.Println(root)
}

func main() {
	root := &Tree{Name: "tree01"}
	left1 := &Tree{Name: "tree02"}
	right1 := &Tree{Name: "tree04"}
	left2 := &Tree{Name: "tree03"}
	root.Left = left1
	root.Right = right1
	left1.Left = left2
	trans_before(root)
	trans_middle(root)
	trans_after(root)

}
