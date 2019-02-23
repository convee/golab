package main

import "sync"

type Set struct {
	m map[int]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: make(map[int]bool),
	}
}

func (s *Set) Add(item int) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item int) {
	delete(s.m, item)
}

func (s *Set) Has(item int) bool {
	return s.m[item]
}

func main() {

}
