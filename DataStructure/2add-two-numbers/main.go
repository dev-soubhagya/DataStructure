package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//to got output from the function
func list2int(head *ListNode) (res []int) {
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return
}

//to take input to the function
func int2list(nums []int) (res *ListNode) {
	if len(nums) == 0 {
		return nil
	}
	res = &ListNode{}
	tmp := res
	for _, v := range nums {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return res.Next
}
func main() {
	///
	fmt.Println(list2int(addTwoNumbers(int2list([]int{0}), int2list([]int{0}))))
	fmt.Println(list2int(
		addTwoNumbers(
			int2list([]int{2, 4, 3}),
			int2list([]int{5, 6, 4}))))
	fmt.Println(list2int(
		addTwoNumbers(
			int2list([]int{9, 9, 9, 9, 9, 9, 9}),
			int2list([]int{9, 9, 9, 9}))))
}

//top rated
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return result
}

//1st approach
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	rem := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + rem
		newnode := ListNode{val % 10, nil}
		rem = val / 10
		if head == nil {
			head = &newnode
			tail = &newnode
		} else {
			tail.Next = &newnode
			tail = tail.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := l1.Val + rem
		newnode := ListNode{val % 10, nil}
		rem = val / 10
		if head == nil {
			head = &newnode
			tail = &newnode
		} else {
			tail.Next = &newnode
			tail = tail.Next
		}
		l1 = l1.Next
	}
	for l2 != nil {
		val := l2.Val + rem
		newnode := ListNode{val % 10, nil}
		rem = val / 10
		if head == nil {
			head = &newnode
			tail = &newnode
		} else {
			tail.Next = &newnode
			tail = tail.Next
		}
		l2 = l2.Next
	}
	if rem > 0 {
		newnode := ListNode{rem, nil}
		if head == nil {
			head = &newnode
			tail = &newnode
		} else {
			tail.Next = &newnode
			tail = tail.Next
		}
	}
	return head
}
*/

//2nd approach
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	ret := cur
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			sum % 10,
			nil,
		}
		cur = cur.Next
		sum = sum / 10
	}

	if sum > 0 {
		cur.Next = &ListNode{sum, nil}
	}

	return ret.Next
}
*/
/***
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{Val: num % 10, Next: nil}
		cur = cur.Next
	}
	return head.Next
}
***/

/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next
	}
	return result
}
*/
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy node as the initial head of the resultant linked list
	dummy := &ListNode{}
	// Create a current pointer variable and initialize it to point to the dummy node
	curr := dummy
	// Create a variable to keep track of the carry from the previous addition, initialize it to 0
	carry := 0

	// While either of the input linked lists (l1 and l2) is not nil
	for l1 != nil || l2 != nil {
		// Initialize x and y to 0
		x := 0
		y := 0

		// If l1 is not nil, assign l1.Val to x and move the l1 pointer to the next node
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		// If l2 is not nil, assign l2.Val to y and move the l2 pointer to the next node
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		// Add the values of x, y and carry, and calculate the sum
		sum := x + y + carry
		// Calculate the carry by dividing the sum by 10
		carry = sum / 10
		// The sum is then stored in the next node of the current pointer
		curr.Next = &ListNode{Val: sum % 10}
		// Move the current pointer to the next node
		curr = curr.Next
	}

	// If there is still a carry left, add a new node to the resultant linked list with the value of the carry
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	// Return the next pointer of the dummy node, which is the actual head of the resultant linked list
	return dummy.Next
}
*/
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy, sum := new(ListNode), 0
	for cur := dummy; l1 != nil || l2 != nil || sum != 0; cur = cur.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}
	return dummy.Next
}
*/
