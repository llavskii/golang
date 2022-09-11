package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
	IsEmpty() bool
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	cItem, ok := c.items[key]
	if ok { // if cache contains this key
		c.queue.MoveToFront(cItem)          // move element first in queue
		cItem.Value = cacheItem{key, value} // update value in element
		c.items[key] = cItem                // update value in cache
	} else { // if cache does not contain this key
		item := cacheItem{key, value}          // create new cache item
		newListItem := c.queue.PushFront(item) // set element first in queue
		c.items[key] = newListItem             // add item to cache
	}
	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		delete(c.items, last.Value.(cacheItem).key)
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	cItem, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(cItem)
		return cItem.Value.(cacheItem).value, ok
	}
	return nil, ok
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) IsEmpty() bool {
	return c.queue.Len() == 0 && len(c.items) == 0
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
