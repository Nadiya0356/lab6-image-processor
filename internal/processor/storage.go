package processor

import "sync"

var LeakCache = make(map[string][]byte)
var mu sync.Mutex

type Storage interface {
	Save(key string, data []byte)
}

type MemoryStorage struct{}

func (m *MemoryStorage) Save(key string, data []byte) {
	mu.Lock()
	defer mu.Unlock()

	LeakCache[key] = data
}
