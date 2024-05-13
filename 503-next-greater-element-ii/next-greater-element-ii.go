/*
    approach: next greater element ; process top first, process ith later
    - this is a classic ngr question, but we have a circular array
    - we can let the outter loop run 2*n
        - inorder to get the correct ith index, we will use i%n
    - we can use process top element ngr approach
        - curr element = ith element
        - "am i(curr) your(top) ngr?"
        - if yes, pop the top, process it
            - out[top] will be current value since its the ngr of top
        - keep doing this until we have elements in stack and until curr is still ngr of new top each time
    - finally, push the ith idx to stack, some future ith element will process it

    - question, should we push the circular elements again into the stack ?
        - it works, if you do, nothing breaks
        - however its not necessary because we have already processed the circular items
        - how do we detect whether this ith element is already processed or not ? whether to push into stack or not ?
        - if i > n, it means, that this ith element is from the circular array, we do not have to push it as its already processed

    time = o(2n)
    space = o(n)
*/

func nextGreaterElements(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {out[i] = -1}
    st := []int{}
    flag := false
    for i := 0; i < n; i++ {
        curr := nums[i]
        for len(st) != 0 && curr > nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = curr
        }
        if i < n{
            st = append(st, i)
        }
        if i == n-1 {
            if flag {break}
            i = -1
            flag = true
        }
    }
    return out
}