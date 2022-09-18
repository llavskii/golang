package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var wg sync.WaitGroup

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errCnt int
	ch := make(chan Task, len(tasks))
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for {
				task, ok := <-ch
				if !ok || errCnt == m {
					break // if channel is closed or error count = limit then break loop and goroutine finishes
				}
				if task != nil && task() != nil { // if channel's element is not nil and returns error
					if errCnt != m { // if error count != limit, then increment error counter and continue
						mu.Lock()
						errCnt++
						mu.Unlock()
					}
				}
			}
			wg.Done()
		}()
	}
	for _, t := range tasks {
		ch <- t
	}
	close(ch)
	wg.Wait()
	if errCnt == m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
