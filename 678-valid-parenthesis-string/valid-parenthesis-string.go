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
*/
func checkValidString(s string) bool {
    open := []int{}
    asteriks := []int{}
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            open = append(open, i)
        } else if s[i] == '*' {
            asteriks = append(asteriks, i)
        } else if s[i] == ')' {
            if len(open) > 0 {
                open = open[:len(open)-1]
            } else if len(asteriks) > 0 {
                // use a previously seen asterik as a open paran
                asteriks = asteriks[:len(asteriks)-1]
            } else {
                // we are at a closing paran
                // without a openining paran and
                // without a asterik in the past
                // therefore this string can never be valid
                return false
            }
        }
    }

    // its possible that we have open parans left
    // now we will use remaining asteriks as the closing paran.
    // for the asteriks to be used as a closing paran, its idx must be
    // AFTER opening paran idx. because a paran pair is valid if we have a open and its corresponding
    // closer is after.
    for len(open) > 0 && len(asteriks) > 0 {
        if asteriks[len(asteriks)-1] > open[len(open)-1] {
            open = open[:len(open)-1]
            asteriks = asteriks[:len(asteriks)-1]
            continue
        }
        break
    }
    return len(open) == 0
}