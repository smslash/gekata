package main

import (
	"fmt"
	"gekata/pkg/fetcher"
	"gekata/pkg/input"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
    urls := input.ReadInput()
    cores := runtime.NumCPU()
    goroutines := make([]int32, cores)
    results := make(chan fetcher.Result, len(urls))
	
    tasks := make(chan string, len(urls))
    for _, url := range urls {
        tasks <- url
    }
    close(tasks)

    var wg sync.WaitGroup
    for i := 0; i < cores; i++ {
        wg.Add(1)
        go func(goroutineID int) {
            defer wg.Done()
            for url := range tasks {
                result := fetcher.Fetch(url)
                results <- result
                atomic.AddInt32(&goroutines[goroutineID], 1)
            }
        }(i)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    for res := range results {
        if res.Error != nil {
            fmt.Printf(res.Error.Error() + "\n")
        } else {
            fmt.Printf("%s;%d;%d;%d\n", res.URL, res.Code, res.Size, res.Duration.Milliseconds())
        }
    }

    for i, count := range goroutines {
        fmt.Printf("%d:%d\n", i + 1, count)
    }
}
