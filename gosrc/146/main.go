package main

import (
	"container/list"
	"fmt"
)

type CacheItem struct {
	value int
	ele   *list.Element
}

type LRUCache struct {
	cap      int
	cache    map[int]*CacheItem
	timeline *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity, make(map[int]*CacheItem), list.New()}
}

func (this *LRUCache) Get(key int) int {
	if item, ok := this.cache[key]; ok {
		this.refresh(item)
		return item.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if item, ok := this.cache[key]; ok {
		item.value = value
		this.refresh(item)
	} else {
		if len(this.cache) == this.cap {
			var key = this.timeline.Remove(this.timeline.Front()).(int)
			delete(this.cache, key)
			// fmt.Println("evicts", key)
		}
		var tlele = this.timeline.PushBack(key)
		this.cache[key] = &CacheItem{value, tlele}
	}
}

func (this *LRUCache) refresh(item *CacheItem) {
	var key = this.timeline.Remove(item.ele)
	item.ele = this.timeline.PushBack(key)
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
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
