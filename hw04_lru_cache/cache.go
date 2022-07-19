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
		l.items[key] = l.queue.PushFront(cacheItem{key, value})
		if l.queue.Len() > l.capacity {
			lastInList := l.queue.Back()
			lastValue, ok := lastInList.Value.(cacheItem)
			if ok {
				delete(l.items, lastValue.key)
				l.queue.Remove(lastInList)
			}
		}
	}
	return res
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	if val, ok := l.items[key]; ok {
		l.queue.MoveToFront(val)
		result, ok := val.Value.(cacheItem)
		if ok {
			return result.value, true
		}
	}
	return nil, false
}

func (l *lruCache) Clear() {
	for k := range l.items {
		delete(l.items, k)
	}
}
