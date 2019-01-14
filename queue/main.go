package main

var data []int

type MyQueue struct {
	data    []int
	p_start int
}

func (q *MyQueue) MyQueue() {
	q.data = []int{}
	q.p_start = 0
}

func (q *MyQueue) EnQueue(x int) {
	q.data = append(q.data, x)
	remove()
}
func main() {

}
