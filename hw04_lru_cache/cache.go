package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	res := false
	if val, ok := l.items[key]; ok {
		val.Value = cacheItem{key, value}
		l.queue.MoveToFront(val)
		res = true
	} else {
		newVal := ListItem{cacheItem{key, value}, nil, nil}
		l.items[key] = &newVal
		l.queue.PushFront(newVal)
		if l.queue.Len() > l.capacity {
			last := l.queue.Back()
			delete(l.items, last.Value.(cacheItem).key)
			l.queue.Remove(last)
		}
	}
	return res
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if val, ok := l.items[key]; ok {
		l.queue.MoveToFront(val)
		return val.Value.(cacheItem).value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	for k := range l.items {
		delete(l.items, k)
	}
}
