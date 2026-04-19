package stats

import (
	"sync"
	"testing"
)

func TestIncrementProcessed(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			IncrementProcessed("jpg")
			wg.Done()
		}()
	}

	wg.Wait()

	if GlobalStats["jpg"] != 100 {
		t.Errorf("Expected 100, got %d", GlobalStats["jpg"])
	}
}
