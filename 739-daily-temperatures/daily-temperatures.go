
// we can also get NGR by traversing left->right
// and asking if I ( curr element ) am YOUR ( top of st ) next-greater-right element ?
// if yes, process the top; pop top, calc dist between i and popped idx, updating out arr
// keep doing this as long as we have elements in the stack and I ( curr element ) is top's next-greater-right element
// the processing is inverted; 
// instead of processing for the curr element ( in classic NGR ),
// we are processing for whats on top of the stack
func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    st := []int{}
    out := make([]int, n)
    for i := 0; i < n; i++ {
        curr := temperatures[i]
        for len(st) != 0 && curr > temperatures[st[len(st)-1]] {
            idx := st[len(st)-1]
            st = st[:len(st)-1]
            out[idx] = i-idx
        }
        st = append(st, i)
    }
    return out
}


/*
    approach: classic NGR
    - next greater element to right
    - start from right hand side
    - and for each ith element, we want its next greater element
    - if we want next element to be accessible easily, we can store them in the stack
    - then each ith element will look at top of the st;
        - remove every element that is <= ith element
        - if top is > ith element, we found ith's NGR
    - because we want distances between 2 indicies
        - we will store index in our st
        - these index's are just reference to each element in temperatures array

    time = o(n)
    space = o(n) 
*/
// func dailyTemperatures(temperatures []int) []int {
//     n := len(temperatures)
//     st := []int{}
//     out := make([]int, n)
//     for i := n-1; i >= 0; i-- {
//         curr := temperatures[i]
//         for len(st) != 0 && temperatures[st[len(st)-1]] <= curr {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i] = st[len(st)-1]-i
//         }
//         st = append(st, i)
//     }
//     return out
// }