/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {    
    if head == nil || head.Next == nil {return head}

    
    // split
    mid := mid(head)
    list2 := mid.Next
    mid.Next = nil
    list1 := head

    // recurse & sort
    list1 = sortList(list1)
    list2 = sortList(list2)

    // merge back 2 sorted LL
    return merge2List(list1,list2)
}

func merge2List(list1, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {return nil}
    if list1 == nil && list2 != nil {return list2}
    if list1 != nil && list2 == nil {return list1}
    dummy := &ListNode{Val: 0}
    tail := dummy
    l1, l2 := list1, list2
    for l1 != nil && l2 != nil {
        l1Next := l1.Next
        l2Next := l2.Next
        if l1.Val < l2.Val {
            l1.Next = nil
            tail.Next = l1
            tail = tail.Next
            l1 = l1Next
        } else {
            l2.Next = nil
            tail.Next = l2
            tail = tail.Next
            l2 = l2Next
        }
    }
    if l1 != nil { tail.Next = l1; tail = tail.Next }
    if l2 != nil { tail.Next = l2; tail = tail.Next }
    return dummy.Next
}

func mid(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var prev *ListNode
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    return prev
}

// time = o(n) + o(nlogn) + o(n)
// space = o(n)
// func sortList(head *ListNode) *ListNode {
//     // lol
//     arr := []int{}
//     curr := head
//     for curr != nil {
//         arr = append(arr, curr.Val)
//         curr = curr.Next
//     }
//     sort.Ints(arr)
//     curr = head
//     i := 0
//     for i < len(arr) {
//         curr.Val = arr[i]
//         i++
//         curr = curr.Next
//     }
//     return head
// }
