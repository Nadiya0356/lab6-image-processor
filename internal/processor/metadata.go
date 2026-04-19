package processor

import (
	"fmt"
	"regexp"
	"time"
)

func processImage(workerID int, storage Storage) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d",
		workerID, time.Now().UnixNano())

	matched, _ := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, data)

	if matched {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())
		storage.Save(key, make([]byte, 1024*10))
	}
}

// 👇 ЦЕ КЛЮЧОВА ФУНКЦІЯ ДЛЯ BENCHMARK
func ProcessImageForBenchmark(storage Storage) {
	processImage(1, storage)
}

func RunWorkerPool(count int, storage Storage) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id, storage)
			}
		}(i)
	}

	select {} // щоб не завершувався
}
