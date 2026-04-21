package main

import (
	"container/list"
	"fmt"
)

// entry holds the key, value, and the frequency of the item.
type entry struct {
	key       int
	value     int
	frequency int
}

// LFUCache represents a least frequently used cache.
type LFUCache struct {
	capacity int
	minFreq  int
	cache    map[int]*list.Element
	freqMap  map[int]*list.List
}

// NewLFUCache initializes an LFU cache with a given capacity.
func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		minFreq:  0,
		cache:    make(map[int]*list.Element),
		freqMap:  make(map[int]*list.List),
	}
}

// Get returns the value of the key if it exists in the cache, otherwise -1.
func (c *LFUCache) Get(key int) int {
	if ele, hit := c.cache[key]; hit {
		c.incrementFrequency(ele)
		return ele.Value.(*entry).value
	}
	return -1
}

// Put adds a value to the cache. If the cache is at capacity, it evicts the LFU item.
func (c *LFUCache) Put(key int, value int) {
	if c.capacity <= 0 {
		return
	}

	if ele, hit := c.cache[key]; hit {
		node := ele.Value.(*entry)
		node.value = value
		c.incrementFrequency(ele)
		return
	}

	if len(c.cache) >= c.capacity {
		c.evict()
	}

	newNode := &entry{key: key, value: value, frequency: 1}
	if c.freqMap[1] == nil {
		c.freqMap[1] = list.New()
	}
	ele := c.freqMap[1].PushFront(newNode)
	c.cache[key] = ele
	c.minFreq = 1
}

func (c *LFUCache) incrementFrequency(ele *list.Element) {
	node := ele.Value.(*entry)
	oldFreq := node.frequency
	c.freqMap[oldFreq].Remove(ele)

	// If the old frequency list is empty and it was the minFreq, increment minFreq
	if c.freqMap[oldFreq].Len() == 0 && oldFreq == c.minFreq {
		c.minFreq++
	}

	node.frequency++
	if c.freqMap[node.frequency] == nil {
		c.freqMap[node.frequency] = list.New()
	}
	newEle := c.freqMap[node.frequency].PushFront(node)
	c.cache[node.key] = newEle
}

func (c *LFUCache) evict() {
	list := c.freqMap[c.minFreq]
	ele := list.Back()
	if ele != nil {
		list.Remove(ele)
		node := ele.Value.(*entry)
		delete(c.cache, node.key)
	}
}

func main() {
	cache := NewLFUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Printf("Get(1): %d\n", cache.Get(1)) // returns 1

	cache.Put(3, 3)                          // evicts key 2 (key 1 has freq 2, key 2 has freq 1)
	fmt.Printf("Get(2): %d\n", cache.Get(2)) // returns -1

	fmt.Printf("Get(3): %d\n", cache.Get(3)) // returns 3 (freq 2)

	cache.Put(4, 4)                          // evicts key 1 (both 1 and 3 have freq 2, evict LRU)
	fmt.Printf("Get(1): %d\n", cache.Get(1)) // returns -1
	fmt.Printf("Get(3): %d\n", cache.Get(3)) // returns 3
	fmt.Printf("Get(4): %d\n", cache.Get(4)) // returns 4
}
