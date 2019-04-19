package main

import "fmt"

type Node struct {
	key    int
	value  int
	before *Node
	after  *Node
}

type LRUCache struct {
	cache    map[int]*Node
	count    int
	capacity int
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{key: 0, value: 0, before: nil, after: nil}
	tail := &Node{key: 0, value: 0, before: nil, after: nil}
	head.after = tail
	tail.before = head
	c := LRUCache{
		cache:    make(map[int]*Node),
		count:    0,
		capacity: capacity,
		head:     head,
		tail:     tail,
	}
	return c
}

func (this *LRUCache) moveToHead(node *Node) {
	node.after = this.head.after
	this.head.after.before = node
	this.head.after = node
	node.before = this.head
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		node.before.after = node.after
		node.after.before = node.before
		this.moveToHead(node)
		return node.value
	}
}

func (this *LRUCache) deleteTail() *Node {
	deleteNode := this.tail.before
	if deleteNode == this.head {
		return nil
	}
	deleteNode.before.after = this.tail
	this.tail.before = deleteNode.before
	deleteNode.before = nil
	deleteNode.after = nil
	return deleteNode
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		if this.count == this.capacity {
			//淘汰最后
			deleteNode := this.deleteTail()
			delete(this.cache, deleteNode.key)
			this.count--
		}
		newNode := &Node{
			key:    key,
			value:  value,
			before: nil,
			after:  nil,
		}
		this.moveToHead(newNode)
		this.cache[key] = newNode
		this.count++
	} else {
		node.value = value
		node.before.after = node.after
		node.after.before = node.before
		this.moveToHead(node)
	}

}

func main() {
	cache := Constructor(2)
	fmt.Println(cache.Get(2))
	cache.Put(2, 6)
	fmt.Println(cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(2))
}
