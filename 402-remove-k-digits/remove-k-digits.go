/*
    1st question that comes to mind is;
    do we remove elements from left side or from right side ?

    9875; k = 1; removing 9 is ideal 

    6789; k = 1; removing 9 is ideal again

    91237 ; k = 2
    ans is ; 123
    we removed 9 and 7  ( from both left and right sides )

    so turns out, there is no definitive answer to this

    lets be greedy and attempt to build smaller numbers from the get-go
    
    approach: stack
    - if curr num is smaller than top of stack
        curr = 1; top = 9
    - then we want to drop 9 in favor of 1
    - therefore while we have elements in stack and curr element trumps top element 
        AND we have k left
        - drop the top
        - keep dropping for as long as we can
    - push the current number
    - finally build the final string from st idx 0 -> len(st)

    # edge case when building numbers; leading 0s
    - number = 201
    - k = 1
    - when i = 0; we push 2 to stack
    - stack = [2]
    - i = 2; number is 0, 0 is smaller, therefore trumps previous number
    - pop the top and now our stack is EMPTY ; []
    - lets push 0 (current number) to stack
    - stack = [0]
    - i is at last digit 1; we have no k left; therefore push 1 to stack
    - stack = [0, 1]
    - now when we building the final string; it will be; "01"
    - "01" is not a valid number
    - because of leading 0s
    - 2 ways we can handle this
    - if current number we are building from stack is 0 and we have nothing in our string; skip this digit
    - or do not push to stack if digit is 0 and stack is empty


    # egde case with monotonic stacks; try increasing and decreasing inputs
    num = 1234
    k = 2
    in stack = [1,2,3,4] <- top
    - we will have all numbers, basically no current ith element trumped previous number
    - i.e no ith element was the nsr of top
    - therefore nothing got popped
    - in this case which numbers do we want to drop
    - well since we know by looking at the stack , that from the top, numbers are in desc
    - therefore top will be the greatest
    - so while we have k left over, drop the top of the stack
    - if stack becomes empty, we need to return "0"



*/
func removeKdigits(num string, k int) string {
    st := []byte{} // chars
    n := len(num)
    for i := 0; i < n; i++ {
        curr := num[i]
        // am i(curr) your(top-of-st) nsr ?
        // if yes, pop it if we still have k
        for len(st) != 0 && curr < st[len(st)-1] && k != 0 {
            st = st[:len(st)-1]
            k--
        }

        if curr == '0' && len(st) == 0 {continue}
        st = append(st, curr)
    } 

    for k != 0 && len(st) != 0 {
        st = st[:len(st)-1]
        k--
    }
    if len(st) == 0 {return "0"}

    out := new(strings.Builder)
    for i := 0; i < len(st); i++ {
        out.WriteByte(st[i])
    }
    return out.String()
}
