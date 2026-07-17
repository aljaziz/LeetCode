/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	hashMap := map[*ListNode]bool{}
	for head != nil {
		if hashMap[head] {
			return true
		}
		hashMap[head] = true
		head = head.Next
	}
	return false
}