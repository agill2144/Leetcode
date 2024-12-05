/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
    approach: brute force ( does not work )
    - iterate thru both lists and create 2 int numbers
    - "digits are stored in reverse order"
    - 342+465 = 807
    - therefore reverse each list first 
    - now write the numbers from the back of res
        - we cant get the digits from left to right unless we convert to str
        - if we can, great ^ 
        - otherwise, we will write 8, then 0 and then 7
        - we can get last digit by: res % 10;
        - and remainng res will be; res / 10;
        - we keep doing this until our res becomes 0
    - the problem? 
        - numbers are too big
        - some integers are so large that they dont fit int64
        - cant use float64, because golang does not support % on float64
    
*/  
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    p1 := l1
    p2 := l2
    head := &ListNode{Val: 0}
    tail := head
    for p1 != nil || p2 != nil {
        p1Val := 0
        if p1 != nil {p1Val = p1.Val; p1 = p1.Next}
        p2Val := 0
        if p2 != nil {p2Val = p2.Val; p2 = p2.Next}
        sum := p1Val + p2Val + carry
        tail.Next = &ListNode{Val: sum % 10}
        tail = tail.Next
        carry = sum / 10
    }
    if carry != 0 {
        tail.Next = &ListNode{Val: carry}
        tail = tail.Next
        carry = 0
    }
    return head.Next
}
