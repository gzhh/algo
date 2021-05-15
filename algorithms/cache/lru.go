package cache

import "container/list"

type LRUCache struct {
	capacity int
	list     *list.List
	cache    map[string]*list.Element
}

type Pair struct {
	key   string
	value string
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		list:     list.New(),
		cache:    make(map[string]*list.Element),
	}
}

func (l *LRUCache) Get(key string) string {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return "none"
}

func (l *LRUCache) Put(key string, value string) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if l.list.Len() >= l.capacity {
			delete(l.cache, l.list.Back().Value.(Pair).key)
			l.list.Remove(l.list.Back())
		}
		l.list.PushFront(Pair{key, value})
		l.cache[key] = l.list.Front()
	}
}
