// Problem:    LRU Cache
// Difficulty: Medium
// Link:       https://neetcode.io/problems/lru-cache
// Source:     Data Structures & Algorithms/lru-cache/submission-0.go (the solution below is a copy of it)
//
// THE PROBLEM
// Build a cache with a fixed capacity:
//
//	Get(key)         the value of the key, or -1 if it is not in the cache
//	Put(key, value)  adds or updates the key. If the cache is then over its
//	                 capacity, the LEAST RECENTLY USED key is thrown out.
//
// Both Get and Put count as "using" a key, and both must run in O(1).
//
// Example:
//
//	Input:  LRUCache(2), Put(1, 10), Get(1), Put(2, 20), Put(3, 30), Get(2), Get(1)
//	Output: -, -, 10, -, -, 20, -1
//
// How to run (from the "local" folder):
//
//	go run ./051-lru-cache
//
// main() is at the bottom of the solution. It runs the test cases and prints
// PASS or FAIL for each one. Add your own cases there to experiment.
package main

import (
	"fmt"
	"os"
)

// ===================== Solution =====================

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

// ===================== Test cases =====================

func main() {
	// The cache remembers what happened before, so the lines starting with
	// "--" show every Put in the output. (A Get also counts as "using" a key.)

	// Example from the problem: a cache with room for 2 keys.
	cache := Constructor(2)
	cache.Put(1, 10)
	fmt.Println(`-- cache := Constructor(2); cache.Put(1, 10)`)
	check("get a key we just put", `cache.Get(1)`,
		cache.Get(1), 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	fmt.Println(`-- cache.Put(2, 20); cache.Put(3, 30)   (the cache is full: key 1 is thrown out)`)
	check("key 2 is still there", `cache.Get(2)`,
		cache.Get(2), 20)
	check("key 1 was thrown out", `cache.Get(1)`,
		cache.Get(1), -1)
	check("key 3 is there", `cache.Get(3)`,
		cache.Get(3), 30)

	// A Get makes a key "recently used", so the OTHER key is thrown out.
	second := Constructor(2)
	second.Put(1, 1)
	second.Put(2, 2)
	fmt.Println(`-- second := Constructor(2); second.Put(1, 1); second.Put(2, 2)`)
	check("use key 1", `second.Get(1)`,
		second.Get(1), 1)
	second.Put(3, 3)
	fmt.Println(`-- second.Put(3, 3)   (key 2 is the least recently used now)`)
	check("key 2 was thrown out", `second.Get(2)`,
		second.Get(2), -1)
	second.Put(4, 4)
	fmt.Println(`-- second.Put(4, 4)   (key 1 is the least recently used now)`)
	check("key 1 was thrown out", `second.Get(1)`,
		second.Get(1), -1)
	check("key 3 is there", `second.Get(3)`,
		second.Get(3), 3)
	check("key 4 is there", `second.Get(4)`,
		second.Get(4), 4)

	// Putting an existing key changes its value AND makes it recently used.
	third := Constructor(2)
	third.Put(1, 1)
	third.Put(2, 2)
	third.Put(1, 100)
	third.Put(3, 3)
	fmt.Println(`-- third := Constructor(2); third.Put(1, 1); third.Put(2, 2); third.Put(1, 100); third.Put(3, 3)`)
	check("updated value", `third.Get(1)`,
		third.Get(1), 100)
	check("key 2 was thrown out, not key 1", `third.Get(2)`,
		third.Get(2), -1)

	// A cache with room for only one key.
	tiny := Constructor(1)
	tiny.Put(1, 1)
	tiny.Put(2, 2)
	fmt.Println(`-- tiny := Constructor(1); tiny.Put(1, 1); tiny.Put(2, 2)`)
	check("capacity 1: old key is gone", `tiny.Get(1)`,
		tiny.Get(1), -1)
	check("capacity 1: new key is there", `tiny.Get(2)`,
		tiny.Get(2), 2)

	// A key that was never put.
	check("unknown key", `tiny.Get(99)`,
		tiny.Get(99), -1)

	report()
}

// ===================== Helpers for the tests =====================

// passed and failed count the test results.
var passed, failed int

// check runs one test case. It prints the input, the result we got and the
// result we expected, with PASS or FAIL in front.
//
//	name:  a short description of the case
//	input: the call we are testing, written as text so it can be printed
//	got:   the answer the solution returned
//	want:  the answer we expect
func check(name string, input string, got interface{}, want interface{}) {
	status := "PASS"
	if fmt.Sprint(got) == fmt.Sprint(want) {
		passed++
	} else {
		failed++
		status = "FAIL"
	}
	fmt.Printf("%s  %s\n", status, name)
	fmt.Printf("      input:    %s\n", input)
	fmt.Printf("      result:   %s\n", show(got))
	fmt.Printf("      expected: %s\n\n", show(want))
}

// show turns a value into text. Strings get quotes, so an empty string is visible.
func show(value interface{}) string {
	if text, ok := value.(string); ok {
		return fmt.Sprintf("%q", text)
	}
	return fmt.Sprint(value)
}

// report prints the total. The program exits with an error if a test failed.
func report() {
	fmt.Printf("%d passed, %d failed\n", passed, failed)
	if failed > 0 {
		os.Exit(1)
	}
}
