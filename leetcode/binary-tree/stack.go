package main

type StackNode struct {
	data [1000]*TreeNode
	pos  int
}

func NewStackNode() *StackNode {
	return &StackNode{}
}
func (s *StackNode) IsEmpty() bool {
	return s.pos == 0
}
func (s *StackNode) Push(n *TreeNode) {
	s.data[s.pos] = n
	s.pos++
}
func (s *StackNode) Pop() (n *TreeNode) {
	s.pos--
	return s.data[s.pos]
}
