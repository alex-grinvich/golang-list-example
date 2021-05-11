package main

import "fmt"


type MyLinkedList struct {
	val  *int
	next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{val: nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {

	for i := 1; i <= index; i++ {
		if this.next !=nil {
			this = this.next
		}else{
			return -1
		}

	}
	return *this.val
}

///** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.val == nil {
		this.val = &val
	} else {
		vals := *this.val
		this.val = &val
		this.next = &MyLinkedList{val: &vals, next: this.next}
	}
}

//
//
///** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.val == nil {
		this.val = &val
	}else if this.next == nil {
		this.next = &MyLinkedList{val:&val}
	} else{
		findElement := true
		for findElement {
			if this.next != nil {
				this = this.next
			} else {
				findElement = false
			}
		}
		tmp := MyLinkedList{val: &val}
		this.next = &tmp
	}

}

//
//
///** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		tmpVal := this.val
		tmpNext := this.next
		this.val = &val
		this.next = &MyLinkedList{val: tmpVal,next: tmpNext}
	}else{
		current := this
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		tmp := MyLinkedList{val: &val, next: current.next}
		current.next = &tmp
	}

}

//
//
///** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		if this.next == nil {
			this.val = nil
		}else{
			tmp:= this.next
			this.val = tmp.val
			this.next = tmp.next
		}
	}else{
		current := this
		pred := this
		for i := 0; i < index; i++ {
			pred = current
			if current.next == nil{
				return
			}
			current = current.next
		}
		pred.next = current.next
	}

}
func main() {

//["MyLinkedList","addAtHead","addAtTail","addAtHead","addAtTail","addAtHead","addAtHead","get","addAtHead","get","get","addAtTail"]
//	[[],[7],[7],[9],[8],[6],[0],[5],[0],[2],[5],[4]]
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(2)
	myLinkedList.AddAtIndex(0,1)
	fmt.Printf("%v", myLinkedList.Get(1))
}
