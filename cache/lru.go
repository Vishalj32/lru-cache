package lru

import (
	"container/list"
	"fmt"
)

//LRUCache definition for Least Recently Used Cache implementation.
type LRUCache struct {
	capacity int                   //defines a cache object of the specified capacity.
	list     *list.List            //DoublyLinkedList for backing the cache value.
	elements map[int]*list.Element //Map to store list pointer of cache mapped to key
}

//KeyPair: defines the cache structure to be stored in LRUCache
type KeyPair struct {
	key   int
	value int
}

//New: creates a new LRUCache object with the defined capacity
func New(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     new(list.List),
		elements: make(map[int]*list.Element, capacity),
	}
}

//Get: returns the cache value stored for the key, also moves the list pointer to front of the list
func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.elements[key]; ok {
		value := node.Value.(*list.Element).Value.(KeyPair).value
		cache.list.MoveToFront(node)
		return value
	}
	return -1
}

//Put: Inserts the key,value pair in LRUCache.
//If list capacity is full, entry at the last index of the list is deleted before insertion.
func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.elements[key]; ok {
		cache.list.MoveToFront(node)
		node.Value.(*list.Element).Value = KeyPair{key: key, value: value}
	} else {
		if cache.list.Len() == cache.capacity {
			idx := cache.list.Back().Value.(*list.Element).Value.(KeyPair).key
			delete(cache.elements, idx)
			cache.list.Remove(cache.list.Back())
		}
	}

	node := &list.Element{
		Value: KeyPair{
			key:   key,
			value: value,
		},
	}

	pointer := cache.list.PushFront(node)
	cache.elements[key] = pointer
}

func (cache *LRUCache) Print() {
	for key, value := range cache.elements {
		fmt.Printf("Key:%d,Value:%+v\n", key, value.Value.(*list.Element).Value.(KeyPair).value)
	}
}

//Keys: returns all the keys present in LRUCache
func (cache *LRUCache) Keys() []interface{} {
	var keys []interface{}
	for k := range cache.elements {
		keys = append(keys, k)
	}
	return keys
}

func (cache *LRUCache) RecentlyUsed() interface{} {
	return cache.list.Front().Value.(*list.Element).Value.(KeyPair).value
}

//Remove: removes the entry for the respective key
func (cache *LRUCache) Remove(key int) {
	if node, ok := cache.elements[key]; ok {
		delete(cache.elements, key)
		cache.list.Remove(node)
	}
}

//Purge: clears LRUCache
func (cache *LRUCache) Purge() {
	cache.capacity = 0
	cache.elements = nil
	cache.list = nil
}
