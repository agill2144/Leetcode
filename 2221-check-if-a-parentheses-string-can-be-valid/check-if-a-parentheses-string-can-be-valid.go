/*
    - anything balancing parans ? 
    - consider stack approach ( from classic balance parans question )
    - why stack?
    - when we run into a closing paran, we want to LOOK back
    - and see if we've had a closing paran?
    - if yes, it means we can balance out the this closing paran
    - hence the need for stack, the need to look back for open parans
    - therefore our stack will track open parans idxs

    approach: 2 stacks
    - similar to LC #678
    - instead of asteriks, we have locked / unlocked chars
    - locked[i] = 0 tells us s[i] can be changed to whatever
    - locked[i] = 1 tell us s[i] cannot be changed at all
    - locked[i] = 0 == asterik
    - EXCEPT we cannot drop unused unlocked idxs as we could with unused asteriks
    - if we have an odd number of unlocked idxs at the end
        - NO MATTER WHAT WE DO, we can never make turn them in balanc parans
        - because balance parans require even number of chars 
    - so if we have even number of unlocked idxs at the end
        - WE CAN TURN THEM INTO BALANCE PARANS
        - no matter how many even parans we have
        - even number of unlocked idxs WILL ALWAYS work!
    - maintain 2 stacks
    - one for tracking idx positions of open parans
    - and one for tracking idx of unlocked positions
    - if current char is unlocked, ( irrespective of paran type )
        - this is unlocked char
        - we can change this to whatever we want
        - therefore add idx to our unlocked st
    - if we run into a open paran, add its idx to open st
    - if we run into a closing paran
        - check if open paran that can be used ?
        - if not, check if we have any unlocked idx ?
        - if not, return false
        - it means, we have too many closing paran
        - and not enough unlocked positions / open parans    
    - everytime we have stack, consider the edge case, what if stack is not empty at the end ?
    - we have 2 stacks, 1 open and 1 for unlocked
    - meaning we have idx positions for extra open parans
    - these open parans can only be closed IF
        - we have unlocked idx positions that occur AFTER open paran idx positions
        - inorder to close a open paran
        - we need a closing paran on right hand side of open paran
        - i.e closing paran idx should be greater than open paran
        - and we can change unlocked idxs to closing paran
        - hence, idx positions of unlocked chars must be AFTER open paran idx positions
    - while we have open parans, we have unlocked idxs, and top
         unlocked idx is after top of open paran
        - pop from both
    - if we still have open parans left after processing the 2nd time
    - it means we have too many open parans
    - and this is invalid
    - if our open st is empty, but we have items in unlocked stack?
    - num of items in unlocked stack should be odd for us to never be able to make
        balanced parans
    - therefore we return true IF len(open) == 0 && len(unlocked) % 2 == 0 
    tc = o(2n)
    sc = o(2n)
*/
func canBeValid(s string, locked string) bool {
    unlocked := []int{}
    open := []int{}
    for i := 0; i < len(s); i++ {
        if locked[i] == '0' {
            unlocked = append(unlocked, i)
        } else if s[i] == '(' {
            open = append(open,  i)
        } else if s[i] == ')' {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(unlocked) > 0{
                unlocked = unlocked[:len(unlocked)-1]
            } else {
                return false
            }
        }
    }
    for len(open) > 0 && len(unlocked) > 0 &&
        unlocked[len(unlocked)-1] > open[len(open)-1] {
            unlocked = unlocked[:len(unlocked)-1]
            open = open[:len(open)-1]
        }

    return len(open) == 0 && len(unlocked) % 2 == 0
}