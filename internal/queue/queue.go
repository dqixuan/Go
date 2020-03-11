package queue

import (
	"fmt"
)

type Queue struct {
	MaxSize int
	Arr [5]int
	Front int
	Rear int
}

func (que *Queue) Add(num int) (err error) {
	if que.Rear + 1 < que.MaxSize {
		que.Rear ++
		que.Arr[que.Rear] = num
		fmt.Println("add success.")
		return
	}
	fmt.Println("array is full.")
	return
}

func (que *Queue) Get() (val int, err error) {
	if que.Front < que.Rear {
		que.Front++
		val = que.Arr[que.Front]
	}
	fmt.Println("array is empty.")
	return
}

func (que *Queue) ShowQueue() {
	for i := que.Front + 1; i <= que.Rear; i++ {
		fmt.Printf("array[%d]=%d", i, que.Arr[i])
	}
	fmt.Println()
}