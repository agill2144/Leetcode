/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
    if head == nil {return 0}
    numsSet := map[int]bool{}
    for i := 0; i < len(nums); i++ {numsSet[nums[i]] = true}

    curr := head
    count := 0
    for curr != nil {
        found := false
        for curr != nil && numsSet[curr.Val] {
            found = true
            curr = curr.Next
        }
        if found {
            count++
        }
        if curr == nil {break}
        curr = curr.Next
    }
    return count
}