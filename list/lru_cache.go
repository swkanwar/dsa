package main

import (
	"container/list"
	"fmt"
)

// LRUCache represents a least recently used cache.
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	ll       *list.List
}

type entry struct {
	key   int
	value int
}

// NewLRUCache initializes an LRU cache with a given capacity.
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		ll:       list.New(),
	}
}

// Get returns the value of the key if it exists in the cache, otherwise -1.
func (c *LRUCache) Get(key int) int {
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

// Put adds a value to the cache. If the cache is at capacity, it evicts the least recently used item.
func (c *LRUCache) Put(key int, value int) {
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}

	ele := c.ll.PushFront(&entry{key, value})
	c.cache[key] = ele

	if c.ll.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
	}
}

func main() {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("Get(1): %d\n", cache.Get(1)) // returns 1

	cache.Put(3, 3)                          // evicts key 2
	fmt.Printf("Get(2): %d\n", cache.Get(2)) // returns -1 (not found)

	cache.Put(4, 4)                          // evicts key 1
	fmt.Printf("Get(1): %d\n", cache.Get(1)) // returns -1 (not found)
	fmt.Printf("Get(3): %d\n", cache.Get(3)) // returns 3
	fmt.Printf("Get(4): %d\n", cache.Get(4)) // returns 4
}
