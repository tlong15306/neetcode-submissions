/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prevNode *ListNode = nil
	var currentNode *ListNode = head
	var nextNode *ListNode

	for currentNode != nil {
		nextNode = currentNode.Next // Lưu lại node tiếp theo
		currentNode.Next = prevNode // Đảo chiều mũi tên của node hiện tại về trước
		prevNode = currentNode      // Dịch prevNode lên một bước
		currentNode = nextNode      // Dịch currentNode lên một bước
	}
	return prevNode
}
