package lru

import (
	"container/list"
	"log"
	"testing"
)

type LruCache struct {
	LruSize int
	Map     map[int]*list.Element
	List    *list.List
}
type Node struct {
	Key int
	val int
}

func NewLruCache(LruSize int) *LruCache {
	return &LruCache{
		LruSize: LruSize,
		Map:     make(map[int]*list.Element),
		List:    list.New(),
	}
}

func (this *LruCache) Put(key int, value int) {
	if val, ok := this.Map[key]; ok {
		val.Value.(*Node).val = value
		this.List.MoveToFront(val)
	} else {
		node := Node{Key: key, val: value}
		front := this.List.PushFront(&node)
		this.Map[key] = front
	}
	if this.List.Len() > this.LruSize {
		element := this.List.Back()
		this.List.Remove(element)
		delete(this.Map, element.Value.(*Node).Key)
	}
}

func (this *LruCache) Get(key int) interface{} {
	if val, ok := this.Map[key]; ok {
		this.List.MoveToFront(val)
		return val.Value.(*Node).val
	}
	return -1
}

func Test_list(t *testing.T) {
	lruCache := NewLruCache(2)
	lruCache.Put(2, 1)
	lruCache.Put(1, 1)
	lruCache.Put(2, 3)
	log.Println(lruCache.Get(2))
	lruCache.Put(4, 1)
}
