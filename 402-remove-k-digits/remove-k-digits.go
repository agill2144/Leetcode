/*
    approach: greedy / stack
    - num = 29
        - if we remove 2; we get 9
        - if we remove 9; we get 2
        - we want 2!
    - num = 92
        - if we remove 9, we get 2
        - if we remove 2, we get 9
        - we want 2!
    
    - we will build a new string from left to right inside a stack
    - and if ith element is the nsr for the top of the stack
    - that is; if ith element is smaller than top of the stack
    - then it does not make sense to keep the top element, we should replace it with the ith (smaller) element
    
    - if we have num = 92
    - we push 9 to stack, since we have nothing to compare yet
    - then i goes to 2; and we have 9 on top
    - does it make sense to continue with 9 and 2 ? no
    - we want the smallest possible digit
    - therefore while ith element is <= top of stack, remove top of stack
    - ith element removes top of stack if ith element is smaller or equal to top of stack and if have k left!

    edge case #1
    - what happens if all digits are increasing order?
    - 1234 ; k = 2
    - then we will have all digits in our stack, and still have k left over
    - then while k != 0, we want to remove the highest digits first ( i,e top of the stack )

    edge case #2
    - when dealing with number formations, dont forget about leading 0's
    - num = 10; k = 1
    - we push 1 to stack
    - and i goes to 0, 0 is smaller, therefore pops the top; k becomes 0
    - stack becomes empty, now do we push 0 ?
    - remember that when we build a final string, we start from idx 0 of stack
    - if push 0 to the stack, our formation will start from 0, which is invalid
    - leading 0 problem
    - either we dont push 0 to stack if stack is empty
    - or we handle leading 0 while building the final string
        - if i == 0 and st[i] == 0; skip adding this char to final string

    time = o(n)
    space = o(n)
*/
func removeKdigits(num string, k int) string {
    st := []byte{}
    for i := 0; i < len(num); i++ {
        curr := num[i]
        // if i am smaller than top element and we still have k's, pop the top
        // i am a better/smaller number to use compared to top of the stack
        // THIS IS HOW WE GET RID OF BIG NUMBERS FROM THE MIDDLE
        // by taking curr number and looking back at what we have had so far
        for len(st) != 0 && curr < st[len(st)-1] && k != 0 {
            st = st[:len(st)-1]
            k--
        }

        if len(st) == 0 && curr == '0' {continue}

        st = append(st, curr)
    }

    // if our digits were are in increasing order
    // then in that case, no ith digit trumped/destroyed previous digits
    // i.e no ith digit was <= top digit
    // ex; 1234
    for k != 0 && len(st) != 0 {
        st = st[:len(st)-1]
        k--
    }

    // build the final string
    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    outStr := out.String()
    if outStr == "" {return "0"}
    return outStr
}