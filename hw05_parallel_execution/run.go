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
	errCnt := 0
	ch := make(chan Task, len(tasks))
	mu := sync.Mutex{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			for {
				mu.Lock()
				if errCnt == m { // if error count = m, then exit
					break
				}
				n, ok := <-ch
				if !ok {
					break // if channel is closed - break loop and goroutine finishes
				}
				if n != nil && n() != nil { // if channel's element is not nil and returns error
					errCnt++
					if errCnt == m { // if error count = m, then exit
						break
					}
				}
				mu.Unlock()
			}
			mu.Unlock()
			wg.Done()
		}(i)
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
