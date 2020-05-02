package main

import "fmt"

func main() {
	fmt.Println("Hello algorithm")
	l1:=new(ListNode)
	fmt.Println(l1.Val)
}

//https://leetcode-cn.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	
	dic := make(map[int]int)
	ret := make([]int, 2)
	
	for index, value := range nums {
		dicValue, ok := dic[value]
		if ok {
			ret[0] = dicValue
			ret[1] = index
			break
		} else {
			dic[target-value] = index
		}
	}
	return ret
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	val1 := 0
	val2 := 0
	addtion := 0
	sum := 0
	newVal := 0
	ret := &ListNode{}
	tmp := ret
	
	for l1 != nil || l2 != nil {
		
		if l1 != nil {
			val1 = l1.Val
		} else {
			val1 = 0
		}
		
		if l2 != nil {
			val2 = l2.Val
		} else {
			val2 = 0
		}
		sum = val1 + val2 + addtion
		newVal = sum % 10
		addtion = sum / 10
		tmp.Val = newVal
		
		if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) || addtion != 0 {
			tmp.Next = &ListNode{addtion, nil}
			tmp = tmp.Next
		}
		
		if l1 != nil {
			l1 = l1.Next
		}
		
		if l2 != nil {
			l2 = l2.Next
		}
		
	}
	
	return ret
}
