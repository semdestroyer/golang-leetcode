package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	ln5 := ListNode{5, nil}
	ln4 := ListNode{4, &ln5}
	ln3 := ListNode{3, &ln4}
	ln2 := ListNode{2, &ln3}
	ln1 := ListNode{1, &ln2}

	listNodeParse(&ln1)
	listNodeParse(removeNthFromEnd(&ln1, 2))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	listStack := make([]*ListNode, 0)
	for head != nil {
		listStack = append(listStack, head)
		head = head.Next
	}
	if len(listStack) == 1 {
		return nil
	}
	if len(listStack) == 2 {
		if n == 1 {
			listStack[0].Next = nil
			return listStack[0]
		}
		if n == 2 {
			return listStack[1]
		}
	}

	if len(listStack) == n {
		return listStack[1]
	}
	listStack[len(listStack)-1-n].Next = listStack[len(listStack)-1-n].Next.Next
	return listStack[0]
}

func listNodeParse(head *ListNode) {
	if head.Next == nil {
		fmt.Println(head.Val)
		return
	}
	listNodeParse(head.Next)
	fmt.Println(head.Val)
}

//TODO: подумать можно ли сделать более "чистое" решение
