/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    set := map[int]struct{}{}
    for i := 0; i < len(nums); i++ {
        set[nums[i]] = struct{}{}
    }
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        _, ok := set[curr.Val]
        if ok {
            if prev != nil {
                prev.Next = next
                curr.Next = nil
                prev = prev
                curr = next
            } else {
                head = next
                curr.Next = nil
                curr = next
            }
            continue
        }
        prev = curr
        curr = next
    }
    return head
}