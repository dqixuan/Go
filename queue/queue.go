package main

import (
	"fmt"
	"os"
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

func main() {
	que := &Queue{
		MaxSize: 5,
		Arr:     [5]int{},
		Front:   -1,
		Rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1. add 添加操作")
		fmt.Println("2. get 获取操作")
		fmt.Println("3. show 显示操作")
		fmt.Println("4. exit 退出操作")
		fmt.Scan(&key)
		switch key {
		case "add":
			fmt.Println("输入一个数字：")
			fmt.Scan(&val)
			que.Add(val)
			fmt.Println("添加成功")
		case "get":
			fmt.Println("获取一个数字：")
			fmt.Println(que.Get())
			fmt.Println("获取成功")
		case "show":
			que.ShowQueue()
			fmt.Println("显示出成功")
		case "exit":
			os.Exit(-1)
		}
	}
}
