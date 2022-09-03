package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
	//"fmt"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// Place your code here.
	wg := new(sync.WaitGroup)
	taskCh := make(chan Task)
	errCount := int32(0)
	wg.Add(n)
	for i:= 0; i < n; i++ {
		go func() error {
			defer wg.Done()
			for task := range taskCh {
				res := task()
				if res != nil { 
					atomic.AddInt32(&errCount, 1)
				}
			}
			return nil
		}() 
	}

	for _, task := range tasks {
		if (atomic.LoadInt32(&errCount) < int32(m)) {
			taskCh <- task
		}
	}
	close(taskCh)
	wg.Wait()
	if (errCount >= int32(m)) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
