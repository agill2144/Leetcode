// process top of stack
func nextGreaterElement(nums1 []int, nums2 []int) []int {
    n1 := len(nums1)
    n2 := len(nums2)
    st := []int{} // idx
    idx := map[int]int{}
    ngr := make([]int, n2)
    for i := 0; i < n2; i++ {
        ngr[i] = -1
        curr := nums2[i]
        idx[curr] = i

        // Am i(curr element) your(top-of-stack; last-in) next-greater-on-right?
        // if yes, pop the top and process it
        // processing top in this case just means updating ngr val in ngr arr
        for len(st) != 0 && curr > nums2[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            ngr[top] = nums2[i]
        }

        // push current element to be processed later
        st = append(st, i)
    }

    // if there are still elements left in st
    // we are processing top
    // if an ith element could not process top of stack
    // it means, there was no ngr for that top of stack
    // and it also means, elements in stack do not have a ngr
    // example: [5,4,3,2,1]
    // st = top-> [5,4,3,2,1]
    // ngr = [-1,-1,-1,-1,-1]
    // therefore nothing to process for leftover elements in top of stack
    out := make([]int, n1)
    for i := 0; i < n1; i++ {
        out[i] = ngr[idx[nums1[i]]]
    }
    return out
}

// brute force: nested iterations
// find ngr of nums2 first, while creating idx map
// then loop over nums1, look for its idx in map and find the ngr
// nested loop depends on ith position (stack hint)
// process ith element ngr pattern
// time = o(n)
// space = o(n)
// func nextGreaterElement(nums1 []int, nums2 []int) []int {
//     n1 := len(nums1)
//     n2 := len(nums2)
//     st := []int{} // holds values
//     idxMap := map[int]int{}
//     ngr := make([]int, n2)
//     for i := n2-1; i >= 0; i-- {
//         curr := nums2[i]
//         ngr[i] = -1
//         idxMap[curr] = i
//         for len(st) != 0 && st[len(st)-1] < curr {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             ngr[i] = st[len(st)-1]
//         }
//         st = append(st, curr)
//     }
//     out := make([]int, n1)
//     for i := 0; i < n1; i++ {
//         out[i] = ngr[idxMap[nums1[i]]]
//     }    
//     return out
// }