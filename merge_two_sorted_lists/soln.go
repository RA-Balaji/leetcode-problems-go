package leetcode

type ListNode struct {
    Val int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	resList := result
	for list1 != nil || list2 != nil {
		if list1.Val < list2.Val {
			resList.Next = list1
			list1 = list1.Next
		} else {
			resList.Next = list2
			list2 = list2.Next
		}
        resList = resList.Next
	}
	if list1 == nil {
        resList.Next = list2
    } else {
        resList.Next = list1
    }
    
    return result.Next
}

