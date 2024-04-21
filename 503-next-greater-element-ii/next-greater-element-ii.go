// ngr on a circular array
// processing top of st instead of processing ith element
/*
    Am i(curr val) your(top of st) next greater element?
    - yes
        - process me(top of the st)
        - and keep processing the top as long as you(i) can
    - no
        - if this ith element is from the 2nd loop ( 2nd pass on circular arr )
            - do not push to stack
        - otherwise push currVal and its idx to the stack
            - the items in the stack are the ones that will be processed
            - each ith element gets processed later, and this element is not in stack yet
            - therefore to process this element, it must be in the stack
            - and some later/next element will process it
    time = o(2n)
    space = o(n)
*/
func nextGreaterElements(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {out[i] = -1}
    st := [][]int{} // [ [val, idx] ]
    for i := 0; i < 2*n; i++ {
        curr := nums[i%n]
        for len(st) != 0 && curr > st[len(st)-1][0] {
            top := st[len(st)-1]
            topIdx := top[1]
            st = st[:len(st)-1]
            out[topIdx] = curr
        }
        // only push elements from the first iteration
        if i < n {
            st = append(st, []int{curr, i})
        }
    }
    return out
}