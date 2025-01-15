/*
    approach: 2 stacks
    - asterik can be used as nothing, (,  or )
    - using valid paranthesis stack approach, we can also consider asteriks
    - use stack to keep track of open parans
    - if we run into a asterik, make a note of it as it can be used in future
    - then as soon as we run into a closing paran, we have few choices;
        1. we have open parans in our stack : remove the top and we are good
        2. we do have open parans but we have seen an asterik before;
            - use previously seen asterik as a open paran for this closing paran to make a valid pair.
            - and move on
        3. we do not have a open paran in stack and we do not have any asteriks
            - this string can never be valid, return false!
    
    - its possible we still have open parans left in the stack
    - and if we still have asteriks, we can use these asteriks as their closing pairs
    - HOWEVER! asterik positions matter!
        - eg: "*()("
        - when we get to last idx, we cannot put a closing paran JUST because we have seen a asterik in the past
        - the closing paran must be after the last idx, therefore asterik idx must be after the last idx
    - we kept respecting the asterik position in the first loop, by looking back at a asterik counter
    - now at the end, if we still have asteriks, we need to know their idx positions for them to become useful
    - therefore we will put the asteriks idx positions into a stack of its own
    - now if top of asteriks stack has a idx < top of open paran, this string is not valid
    - but if top of asteriks stack has a idx > top of open paran, pop both and continue processing rest of open parans
    - eg: "(*()*"
    - we successfully balance the middle pair of parans
    - at the end, we are left with 2 asteriks and a 1 open paran
    - then we can check, whether an asterik shows up after open paran idx
        - i.e asteriks[len(asteriks)-1] > open[len(open)-1] {if yes, pop both}
    - the input string is only valid is we have successfully processed all of the open parans!
    - therefore return true if len(open) == 0
    time = o(n)
    space = o(2n)

    
    approach: o(1) space
    - can we simply keep a count of open parans instead of stack?
    - if we run into a closing paran, we can decrement the counter
    - if the counter goes negative, it means, we had too many closing parans
    - hence invalid
    - what about asteriks tho?
    - an asterik can be open paran or closing paran
    - because we are trying to balance out closing parans, i.e 
        checking if we have seen open parans before ( open counter )
        we can use asteriks as open paran in the counter
    - if our open paran goes negative, EVEN AFTER including ASTERIKS in open paran count
    - then we know for sure that we just ran into too many closing parans
    - that using asteriks as open could not help...

    - but what if open count was still > 0 at the end of the loop?
    - how does this guarantee that we did not have too many open parans?
    - how does this guarantee that we do not have extra open parans?
    - the above left->right pass helps us verify we do not have too many closing parans!
    - so then, lets use a right->left pass to help us verify whether we have too many opening parans!
    
    
    - for a closing paran to be balanced, we need a opening paran
    - therefore we can take another pass backward
    - instead of counting open, we now count closing
    - and this time, we treat asteriks as closing paran
    - if we run into a open paran, we decrement closing paran
    - and if closing counter goes negative, 
    - it means we just ran into too many open parans that even
    - after using asteriks as closing, we could not balance it out
    - therefore return false
    
    - at the end;
    - we have verified that we do not have too many close parans
        by using asteriks as open paran
    - we have verified that we do not have too many open parans
        by using asteriks as close paran
    
    - what if our counter for both open and closing > 0 ?
    - thats okay, because one of the choices of asteriks
    - is that we can simply decide to delete that char all together

    tc = o(n)
    sc = o(1)


*/
func checkValidString(s string) bool {
    n := len(s)
    if n == 0 {return true}
    open := 0
    close := 0
    for i := 0; i < n; i++ {
        // handle using asterik as open
        if s[i] == '(' || s[i] == '*' {
            open++
        } else {
            open--
            if open < 0 {return false}
        }

        // simulatenously, traverse from righ to left
        // handle using asterik as close
        if s[n-1-i] == ')' || s[n-1-i] == '*' {
            close++
        } else {
            close--
            if close < 0 {return false}
        }

    }
    return true
}