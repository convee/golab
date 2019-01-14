package main

type MyHashSet struct {
	data []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		data: make([]int, 1000000),
	}
}

func (this *MyHashSet) Add(key int) {
	this.data[key] = 1
}

func (this *MyHashSet) Remove(key int) {
	this.data[key] = 0
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	if this.data[key] == 1 {
		return true
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {

}
