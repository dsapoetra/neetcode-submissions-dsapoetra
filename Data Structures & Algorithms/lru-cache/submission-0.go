// Idea: we need two things to be fast:
//   1) find a key quickly            -> a map from key to node
//   2) know the order of last use    -> a doubly linked list:
//        least recently used ... most recently used
// Every Get or Put moves that node to the "most recent" end.
// When the cache is too full, we remove the node at the "least recent" end.
// Two fake nodes (oldest and newest) sit at the two ends, so we never
// have to handle an empty list as a special case.
// Get and Put are both O(1).

type cacheNode struct {
	key   int
	value int
	prev  *cacheNode
	next  *cacheNode
}

type LRUCache struct {
	capacity int
	nodes    map[int]*cacheNode
	oldest   *cacheNode // fake node before the least recently used node
	newest   *cacheNode // fake node after the most recently used node
}

func Constructor(capacity int) LRUCache {
	oldest := &cacheNode{}
	newest := &cacheNode{}
	oldest.next = newest
	newest.prev = oldest
	return LRUCache{capacity: capacity, nodes: map[int]*cacheNode{}, oldest: oldest, newest: newest}
}

// remove takes a node out of the list.
func (this *LRUCache) remove(node *cacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addNewest puts a node at the "most recently used" end.
func (this *LRUCache) addNewest(node *cacheNode) {
	last := this.newest.prev
	last.next = node
	node.prev = last
	node.next = this.newest
	this.newest.prev = node
}

func (this *LRUCache) Get(key int) int {
	node, found := this.nodes[key]
	if !found {
		return -1
	}
	// it was just used: move it to the newest end
	this.remove(node)
	this.addNewest(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	// the key already exists: take the old node out first
	if node, found := this.nodes[key]; found {
		this.remove(node)
	}

	node := &cacheNode{key: key, value: value}
	this.nodes[key] = node
	this.addNewest(node)

	// too many keys: throw away the least recently used one
	if len(this.nodes) > this.capacity {
		lru := this.oldest.next
		this.remove(lru)
		delete(this.nodes, lru.key)
	}
}
