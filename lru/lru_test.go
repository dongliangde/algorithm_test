package lru

import (
	"log"
	"testing"
)

type LRUCache struct {
	size       int //缓存大小
	initSize   int
	cache      map[int]*DLinkNode //哈希Map
	head, tail *DLinkNode         //链表 链表头，链表尾
}

type DLinkNode struct {
	key, value int
	prev       *DLinkNode //前一个
	next       *DLinkNode //后一个
}

//初始化Node节点
func initDLinkedNode(key, value int) *DLinkNode {
	return &DLinkNode{key: key, value: value}
}

//初始化LRU缓存
func Constructor(initSize int) LRUCache {
	lRUCache := LRUCache{
		initSize: initSize,
		cache:    make(map[int]*DLinkNode),
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	lRUCache.head.next = lRUCache.tail
	lRUCache.tail.next = lRUCache.head
	return lRUCache
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.cache[key]; ok {
		this.moveToHead(val)
		return val.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.initSize {
			removeTail := this.removeTail()
			delete(this.cache, removeTail.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}
func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

//修改当前接点信息
func (this *LRUCache) removeNode(node *DLinkNode) {
	//当前node节点前一个node节点指定的下一个node节点替换为当前的下一个节点
	node.prev.next = node.next
	//当前下一个node的前一个node节点替换为当前node节点的前一个
	node.next.prev = node.prev
}

//添加头
func (this *LRUCache) addToHead(node *DLinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

//移除尾
func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func Test(t *testing.T) {
	lruCache := Constructor(2)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	log.Println(lruCache.Get(1))
	lruCache.Put(3, 3)
	log.Println(lruCache.Get(2))
}
