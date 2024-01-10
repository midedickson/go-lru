package main

import "fmt"

type Node[V comparable] struct {
	Value interface{}
	Next  *Node[V]
	Prev  *Node[V]
}

func newNode[V comparable](value V) *Node[V] {
	node := &Node[V]{Value: value}
	return node
}

type LRUCache[K, V comparable] struct {
	length        int
	capacity      int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
}

func newLRUCache(capacity int) *LRUCache[string, interface{}] {
	lruCache := &LRUCache[string, interface{}]{}
	lruCache.Init(capacity)

	return lruCache
}

func (r *LRUCache[K, V]) Init(capacity int) {
	r.length = 0
	r.capacity = capacity
	r.head = nil
	r.tail = nil
	r.lookup = make(map[K]*Node[V])
	r.reverseLookup = make(map[*Node[V]]K)
}

func (r *LRUCache[K, V]) Update(key K, value V) {
	//does it exist?
	node, ok := r.lookup[key]
	if !ok {
		// if not found, return None
		node = newNode(value)
		r.prepend(node)
		r.length = r.length + 1
		r.trimCache()
		r.lookup[key] = node
		r.reverseLookup[node] = key
	} else {
		r.detach(node)
		r.prepend(node)
		node.Value = value
	}
	//  if it doesn't, we need to insert it
	//     - check capacity and evict lru entries
	//if it does, we need to update to the front of the list
	//and udate the value
}

func (r *LRUCache[K, V]) Get(key K) interface{} {
	// check the cache for existence
	node, ok := r.lookup[key]
	if !ok {
		// if not found, return None
		return nil
	}
	// update the value we found and move it to the front
	r.detach(node)
	r.prepend(node)
	return node.Value

}

func (r *LRUCache[K, V]) trimCache() {
	// trim cache to fit capacity
	if r.length <= r.capacity {
		return
	}
	tail := (r.tail)
	r.detach(tail)
	key := r.reverseLookup[tail]
	delete(r.lookup, key)
	delete(r.reverseLookup, tail)
	r.length = r.length - 1

}

func (r *LRUCache[K, V]) detach(node *Node[V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if node == r.head {
		r.head = r.head.Next
	}
	if node == r.tail {
		r.tail = r.tail.Prev
	}

	node.Prev = nil
	node.Next = nil
}
func (r *LRUCache[K, V]) prepend(node *Node[V]) {
	if r.head == nil {
		r.tail = node
		r.head = node
		return
	}
	node.Next = r.head
	r.head.Prev = node
	r.head = node
}

func main() {
	lru := newLRUCache(3)
	fmt.Printf("Current Length: %d\n", lru.length)

	lru.Update("key1", 1)
	lru.Update("key2", 2)
	lru.Update("key3", 3)
	fmt.Printf("Current Length: %d\n", lru.length)

}
