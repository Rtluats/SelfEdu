//////////////////////////////////////////////////////////////////////
//
// Given is some code to cache key-value pairs from a database into
// the main memory (to reduce access time). Note that golang's map are
// not entirely thread safe. Multiple readers are fine, but multiple
// writers are not. Change the code to make this thread safe.
//

package main

import (
	"container/list"
	"sync"
	"testing"
)

// CacheSize determines how big the cache can grow
const CacheSize = 100

// KeyStoreCacheLoader is an interface for the KeyStoreCache
type KeyStoreCacheLoader interface {
	// Load implements a function where the cache should gets it's content from
	Load(string) string
}

type page struct {
	Key   string
	Value string
}

// KeyStoreCache is a LRU cache for string key-value pairs
type KeyStoreCache struct {
	cache sync.Map
	pages list.List
	load  func(string) string
	mutex sync.Mutex
}

// New creates a new KeyStoreCache
func New(load KeyStoreCacheLoader) *KeyStoreCache {
	return &KeyStoreCache{
		load:  load.Load,
		cache: sync.Map{},
		mutex: sync.Mutex{},
	}
}

func (k *KeyStoreCache) Len() int {
	var length int
	k.cache.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	return length
}

// Get gets the key from cache, loads it from the source if needed
func (k *KeyStoreCache) Get(key string) string {
	e, ok := k.cache.LoadOrStore(
		key,
		&list.Element{
			Value: page{
				Key:   key,
				Value: k.load(key),
			},
		},
	)
	if ok {
		k.pages.MoveToFront(e.(*list.Element))
		return e.(*list.Element).Value.(page).Value
	}

	// Miss - load from database and save it in cache
	p := page{key, e.(*list.Element).Value.(page).Value}

	// if cache is full remove the least used item
	if k.Len() >= CacheSize {
		end := k.pages.Back()
		k.pages.Remove(end)
		// remove from map
		k.cache.Delete(end.Value.(page).Key)

	}
	k.pages.PushFront(p)
	f := k.pages.Front()
	k.cache.Store(key, f)
	return p.Value
}

// Loader implements KeyStoreLoader
type Loader struct {
	DB *MockDB
}

// Load gets the data from the database
func (l *Loader) Load(key string) string {
	val, err := l.DB.Get(key)
	if err != nil {
		panic(err)
	}

	return val
}

func run(t *testing.T) (*KeyStoreCache, *MockDB) {
	loader := Loader{
		DB: GetMockDB(),
	}
	cache := New(&loader)

	RunMockServer(cache, t)

	return cache, loader.DB
}

func main() {
	run(nil)
}
