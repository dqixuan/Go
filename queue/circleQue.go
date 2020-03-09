package queue

import (
	"errors"
	"fmt"
)

type CircleQueue struct {
	MaxSize int // 4
	Arr [4]int
	Head   int
	Tail   int
}

func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue is full")
	}
	// 尾部不含最后一个元素
	this.Arr[this.Tail] = val
	this.Tail++
	return
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val = this.Arr[this.Head]
	this.Head++
	return
}

// 判断队列是否已满
func (this *CircleQueue) IsFull() (isFull bool) {
	return (this.Tail + 1) % this.MaxSize == this.Head
}

// 判断队列是否为空
func (this *CircleQueue) IsEmpty() (isEmpty bool) {
	return this.Head == this.Tail
}

// 返回队列有多少元素
func (this *CircleQueue) Size() (size int ) {
	return (this.Tail + this.MaxSize - this.Head) % this.MaxSize
}

// 显示队列
func (this *CircleQueue) ShowQueue() {
	for i := this.Head; i % this.MaxSize < this.Tail; i++ {
		fmt.Printf("%d\t", this.Arr[i])
	}
}