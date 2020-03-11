package singleLisk

import "fmt"

type SingleLink struct {
	no int
	Name string
	NickName string
	next *SingleLink
}

func InsertAtTail(head *SingleLink, newNode *SingleLink) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
	return
}
// showLink
func ShowLink (head *SingleLink) {
	temp := head
	for {
		if temp == nil {
			break
		}
		temp = temp.next
		if temp != nil {
			fmt.Println(temp.no, temp.Name, temp.NickName)
		}
	}
}
