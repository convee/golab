package main

// 思路：题目要求实现O(1)时间复杂度的获取数据和写入数据。哈希表可以实现O(1)复杂度的获取数据，但是没有办法写入数据，因为没有办法判断大小，就不知道这个数据应该写到哪里。双向链表可以以O(1)时间复杂度，很方便地实现数据的插入和删除，但是没有办法直接定位。所以我们采用哈希表和双向链表相结合的方法。
//1、通过哈希表来定位到双向链表节点的入口：哈希表的（key，val）分别取（key，ListNode*），这个key就是题目中的key；通过双向链表来实现数据的插入和删除，双线链表存的值为int key, int value，就是题目中的（key，val）。
//2、获取数据的时候：
//如果密钥存在于缓存中，那么返回缓存的value值，同时在列表中将该节点删除并且插入到链表的最前端；
//如果密钥不存在于缓存中，返回-1。
//3、写入数据的时候：
//如果密钥存在，在链表中将该结点删除并插入到最前端；
//如果密钥不存在，如果缓存容量达到上限删除链表的最后一个元素，然后将该节点插入到链表的最前端；哈希表中插入该元素。
type LRUCache struct {
	Capacity int
	Count    int
	Length   int
	Size     int
	Caches   []*Cache
}

type Cache struct {
	Key   int
	Value int
	Count int
	Next  *Cache
}

func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		Capacity: capacity,
		Count:    0,
		Length:   0,
		Caches:   make([]*Cache, capacity/10+1),
		Size:     capacity/10 + 1,
	}
	return lc
}

func (this *LRUCache) Get(key int) int {
	cache := this.Caches[key%this.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			this.Count++
			cache.Count = this.Count
			result = cache.Value
			if cache.Next != nil {
				t := cache
				if temp == nil {
					this.Caches[key%this.Size] = cache.Next
				} else {
					temp.Next = cache.Next
				}
				for cache.Next != nil {
					cache = cache.Next
				}
				cache.Next = t
				t.Next = nil
			}
			break
		}
		temp = cache
		cache = cache.Next
	}
	return result
}

func (this *LRUCache) Put(key int, value int) {
	cache := this.Caches[key%this.Size]
	result := -1
	var temp *Cache
	for cache != nil {
		if cache.Key == key {
			this.Count++
			cache.Count = this.Count
			cache.Value = value
			result = value
			if cache.Next != nil {
				t := cache
				if temp == nil {
					this.Caches[key%this.Size] = cache.Next
				} else {
					temp.Next = cache.Next
				}
				for cache.Next != nil {
					cache = cache.Next
				}
				cache.Next = t
				t.Next = nil
			}
			break
		}
		temp = cache
		cache = cache.Next
	}
	if result != -1 {
		return
	}
	for temp != nil && temp.Next != nil {
		temp = temp.Next
	}
	if this.Length == this.Capacity {
		min := this.Count + 1
		pin := 0
		for i := 0; i < this.Size; i++ {
			if this.Caches[i] != nil && this.Caches[i].Count < min {
				min = this.Caches[i].Count
				pin = i
			}
		}
		if temp != nil && this.Caches[pin].Key == temp.Key {
			temp = nil
		}
		this.Caches[pin] = this.Caches[pin].Next
	} else {
		this.Length++
	}
	this.Count++
	t := &Cache{
		Key:   key,
		Value: value,
		Count: this.Count,
	}
	if temp == nil {
		this.Caches[key%this.Size] = t
	} else {
		temp.Next = t
	}
}

func main() {

}
