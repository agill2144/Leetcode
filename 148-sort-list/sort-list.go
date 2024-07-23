/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func sortList(head *ListNode) *ListNode {
    // lol
    arr := []int{}
    curr := head
    for curr != nil {
        arr = append(arr, curr.Val)
        curr = curr.Next
    }
    sort.Ints(arr)
    curr = head
    i := 0
    for i < len(arr) {
        curr.Val = arr[i]
        i++
        curr = curr.Next
    }
    return head
}
