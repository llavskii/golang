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
	if m == 0 {
		return ErrErrorsLimitExceeded
	}
	var errCnt int
	ch := make(chan Task, n)
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			for {
				task, ok := <-ch
				if !ok {
					break // if channel is closed than break loop and goroutine finishes
				}
				mu.Lock()
				if errCnt == m { // if error count = limit than break loop and goroutine finishes
					mu.Unlock()
					break
				}
				mu.Unlock()
				if task != nil && task() != nil { // if channel's element is not nil and returns error
					mu.Lock()
					if errCnt != m { // if error count != limit, than increment error counter and loop continues
						errCnt++
					}
					mu.Unlock()
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
