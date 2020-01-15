package hw_8

import (
	"errors"
	"fmt"
	"math/rand"
)

const errPercent = 0

func runTask(ch chan error, t func() error) {
	ch <- t()
}

func TaskDoer(tasks []func() error, n int, errMax int) {
	tasksCnt := len(tasks)
	ch := make(chan error, n)
	runningTasksCnt := 0
	doneTasksCnt := 0
	errorsCnt := 0

	for {
		fmt.Printf("Running task %d\n", runningTasksCnt)
		go runTask(ch, tasks[runningTasksCnt])
		runningTasksCnt++
		if runningTasksCnt == n {
			break
		}
	}

	for {
		if <-ch != nil {
			errorsCnt++
		}
		if runningTasksCnt < tasksCnt {
			fmt.Printf("Running task %d\n", runningTasksCnt)
			go runTask(ch, tasks[runningTasksCnt])
			runningTasksCnt++
		}
		if errorsCnt > errMax {
			fmt.Printf("Exit with error! Errors count: %d\n", errorsCnt)
			break
		}
		doneTasksCnt++
		if doneTasksCnt == tasksCnt {
			fmt.Printf("All tasks done! Errors count: %d\n", errorsCnt)
			break
		}
	}
}

func task(i int) func() error {
	return func() error {
		if rand.Intn(100) < errPercent {
			return errors.New("it is error")
		}
		fmt.Printf("Task %d\n", i)
		return nil
	}
}

func Check() {
	n := 3
	errMax := 1
	tasks := []func() error{task(0), task(1), task(2), task(3), task(4), task(5)}
	TaskDoer(tasks, n, errMax)
}
