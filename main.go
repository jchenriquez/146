package main

import "fmt"

type LinkedList struct {
	Prev *LinkedList
	Next *LinkedList
	Key  int
	Val  int
}

type LRUCache struct {
	Head        *LinkedList
	Tail        *LinkedList
	FrequencyMp map[int]*LinkedList
	Capacity    int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{nil, nil, make(map[int]*LinkedList, capacity), capacity}
}

func increaseAccess(ls *LinkedList, cache *LRUCache) {
	if ls.Prev != nil {
		prev := ls.Prev
		next := ls.Next

		prev.Next = next

		if next != nil {
			next.Prev = prev
		}

		if prev.Next == nil {
			cache.Tail = prev
		}

		ls.Next = cache.Head
    ls.Prev = nil
		cache.Head.Prev = ls
		cache.Head = ls
	}
}

func (this *LRUCache) Get(key int) int {
	ls, in := this.FrequencyMp[key]

	if !in {
		return -1
	} else {
		increaseAccess(ls, this)
		return ls.Val
	}
}

func (this *LRUCache) Put(key int, value int) {
	ls, in := this.FrequencyMp[key]

	if !in {
		if len(this.FrequencyMp) >= this.Capacity {
			delete(this.FrequencyMp, this.Tail.Key)
			if this.Tail.Prev == nil {
				this.Head = nil
				this.Tail = nil
			} else {
				tail := this.Tail
				tail.Prev.Next = nil
				this.Tail = tail.Prev
				tail.Prev = nil
			}
		}
		ls := LinkedList{nil, nil, key, value}
		if this.Tail == nil {
			this.Head = &ls
			this.Tail = &ls
		} else {
			ls.Next = this.Head
			this.Head.Prev = &ls
			this.Head = &ls
		}

		this.FrequencyMp[key] = &ls
	} else {
		ls.Val = value
		increaseAccess(ls, this)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Printf("1: %d\n", cache.Get(1))

	cache.Put(3, 3)

	fmt.Printf("2: %d\n", cache.Get(2))

	cache.Put(4, 4)
	fmt.Printf("1: %d\n", cache.Get(1))
	fmt.Printf("3: %d\n", cache.Get(3))
	fmt.Printf("4: %d\n", cache.Get(4))
}
