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
	fmt.Println(root.Name)
	trans_before(root.Left)
	trans_before(root.Right)
}

//中序遍历
func trans_middle(root *Tree) {
	if root == nil {
		return
	}
	trans_middle(root.Left)
	fmt.Println(root.Name)
	trans_middle(root.Right)
}

//后序遍历
func trans_after(root *Tree) {
	if root == nil {
		return
	}
	trans_after(root.Left)
	trans_after(root.Right)
	fmt.Println(root.Name)
}

func main() {
	root := new(Tree)
	root.Name = "tree01"
	left1 := new(Tree)
	left1.Name = "tree02"
	root.Left = left1
	right1 := new(Tree)
	right1.Name = "tree04"
	root.Right = right1
	left2 := new(Tree)
	left2.Name = "tree03"
	left1.Left = left2
	fmt.Println("----------trans_before-----------")

	trans_before(root)
	fmt.Println("----------trans_middle-----------")
	trans_middle(root)
	fmt.Println("----------trans_after-----------")
	trans_after(root)

}
