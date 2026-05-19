/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	fast := head
	slow := head


	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next

	}

	secondHead := slow.Next
	slow.Next = nil 


	var prev *ListNode

	for secondHead != nil {

		temp := secondHead.Next
		secondHead.Next = prev
		prev = secondHead
		secondHead = temp

	}

	first := head
	second := prev

	for second != nil {

		temp1 := first.Next
		temp2 := second.Next

		first.Next = second
		second.Next = temp1

		first = temp1
		second = temp2
	}







	
}
