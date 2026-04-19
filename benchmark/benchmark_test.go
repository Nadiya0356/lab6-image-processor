package benchmark

import (
	"lab6-image-processor/internal/processor"
	"testing"
)

func BenchmarkProcessImage(b *testing.B) {
	storage := &processor.MemoryStorage{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		processor.ProcessImageForBenchmark(storage)
	}
}
