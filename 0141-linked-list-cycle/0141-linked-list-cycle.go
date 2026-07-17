/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
	hashMap := map[*ListNode]bool{}
	for head.Next != nil {
		if hashMap[head] {
			return true
		}
		hashMap[head] = true
		head = head.Next
	}
	return false
}