package lru_cache

import (
	"container/heap"
	"log"
	"time"
)

type Cache interface {
	Set(key int, value interface{})
	Get(key int) interface{}
}

type Item struct {
	key       int
	value     interface{}
	timestamp int64
	index     int
}

type Items []*Item

func (this Items) Len() int {
	return len(this)
}

func (this Items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this Items) Less(i, j int) bool {
	return this[i].timestamp < this[j].timestamp
}

func (this *Items) Push(val interface{}) {
	newItem := val.(*Item)
	*this = append(*this, newItem)
}

func (this *Items) Pop() interface{} {
	old := *this
	oldLen := len(old)
	top := old[oldLen-1]
	*this = old[0 : oldLen-1]
	return top
}

type LRUCache struct {
	lruHeap  Items
	cache    map[int]*Item
	capacity int
}

func New(capacity int) *LRUCache {
	// TODO: Validate capacity
	items := make(Items, 0)
	heap.Init(&items)
	return &LRUCache{lruHeap: items, cache: make(map[int]*Item, capacity), capacity: capacity}
}

func (this *LRUCache) Set(key int, value interface{}) {
	ts := time.Now().UnixNano()
	if item, present := this.cache[key]; present {
		item.value = value
		item.timestamp = ts
		heap.Fix(&this.lruHeap, item.index)
	} else {
		newItem := &Item{key: key, value: value, timestamp: ts}
		if this.capacity == len(this.cache) {
			top := heap.Pop(&this.lruHeap).(*Item)
			log.Printf("Removing from cache: %v", top)
			delete(this.cache, top.key)
		}
		this.cache[key] = newItem
		heap.Push(&this.lruHeap, newItem)
	}
}

func (this *LRUCache) Get(key int) interface{} {
	ts := time.Now().UnixNano()
	if item, present := this.cache[key]; present {
		item.timestamp = ts
		heap.Fix(&this.lruHeap, item.index)
		return item.value
	} else {
		empty := &Item{}
		return empty.value
	}
}
