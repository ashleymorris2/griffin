package blueprint

import (
	"gopkg.in/yaml.v3"
	"os"
	"sync"
)

type Metadata struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	FilePath    string `yaml:"-"`
}

type ReadResult struct {
	Index int
	Item  Metadata
	Err   error
}

func ReadMetadataFromFiles(files []string) chan ReadResult {
	var wg sync.WaitGroup // Counts running goroutines

	results := make(chan ReadResult, len(files)) // Buffered channel to hold the results
	limit := make(chan struct{}, 8)              // Ensures that no more than n goroutines are ran concurrently

	for i, file := range files {
		wg.Add(1)
		go func(i int, file string) {
			defer wg.Done()
			limit <- struct{}{}        // Acquire a slot, blocks if already at the set limit
			defer func() { <-limit }() // Release slot once complete

			data, err := readMetadataFromFile(file)
			if err != nil {
				results <- ReadResult{Index: i, Err: err}
				return
			}

			results <- ReadResult{
				Index: i,
				Item:  data,
			}
		}(i, file)
	}

	wg.Wait() // Waits for all goroutines to complete
	close(results)

	return results
}

func readMetadataFromFile(path string) (Metadata, error) {
	var meta Metadata

	data, err := os.ReadFile(path)
	if err != nil {
		return meta, err
	}

	if err := yaml.Unmarshal(data, &meta); err != nil {
		return meta, err
	}

	meta.FilePath = path
	return meta, nil
}
