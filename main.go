package main

type Node[V comparable] struct {
	Value V
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
	// Check if the key already exists
	node, ok := r.lookup[key]

	if !ok {
		// If not found, create a new node
		node = newNode(value)
		r.prepend(node)
		r.length++
		r.trimCache()
		r.lookup[key] = node
		r.reverseLookup[node] = key
	} else {
		// If the key exists, move the node to the front and update the value
		r.detach(node)
		r.prepend(node)
		node.Value = value
	}
}

func (r *LRUCache[K, V]) Get(key K) (value V) {
	// check the cache for existence
	node, ok := r.lookup[key]
	if !ok {
		// if not found, return zero-value of V
		return
	}
	// update the value we found and move it to the front
	r.detach(node)
	r.prepend(node)
	value = node.Value
	return
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
	r.length--

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
