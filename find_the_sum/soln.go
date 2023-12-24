package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var num1, num2 int
    for l1.Next != nil {
		num1 = num1 * 10 + l1.Val
		l1 = l1.Next
	}
	for l2.Next != nil {
		num2 = num2 * 10 + l1.Val
		l1 = l1.Next
	}
	resSum := num1 + num2
	var resList ListNode
	for resSum > 0 {
		resList.Val = resSum % 10
        resList.Next = &ListNode{
			Val: 0,
			Next: &ListNode{},
		}
		resList = *resList.Next
		resSum = resSum / 10
	}

	return &resList
}

