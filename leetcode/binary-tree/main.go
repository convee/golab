package main

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

func preorderTraversalByStack(root *TreeNode) []int {
	r := make([]int, 0)
	if root == nil {
		return r
	}
	stack := StackNode{}
	stack.Push(root)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		r = append(r, cur.Val)
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}
	}
	return r
}

func preorderNoRec(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	var nodes []int

	stack = append(stack, root)

	for len(stack) != 0 {
		lastNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nodes = append(nodes, lastNode.Val)

		if lastNode.Right != nil {
			stack = append(stack, lastNode.Right)
		}
		if lastNode.Left != nil {
			stack = append(stack, lastNode.Left)
		}
	}

	return nodes
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		var seq []int
		for i := 0; i < l; i++ {
			node := queue[i]
			seq = append(seq, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
		res = append(res, seq)
	}
	return res
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftH := maxDepth(root.Left)
	rightH := maxDepth(root.Right)
	max := max(leftH, rightH)
	return 1 + max
}
func max(left, right int) int {
	if left > right {
		return left
	}
	return right
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
	// fmt.Println(preorderTraversal(root))
	// fmt.Println(levelOrder(root))
	maxDepth(root)
}
