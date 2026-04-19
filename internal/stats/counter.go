package stats

import "sync"

var mu sync.Mutex
var GlobalStats = make(map[string]int)

func IncrementProcessed(imageType string) {
	mu.Lock()
	defer mu.Unlock()
	GlobalStats[imageType]++
}
