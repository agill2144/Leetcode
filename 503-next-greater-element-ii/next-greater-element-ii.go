// processing top pattern ( ideal pattern )
func nextGreaterElements(nums []int) []int {
    n := len(nums)
    st := []int{} // idx
    out := make([]int, n)
    for i := 0; i < n; i++ {out[i] = -1}
    for i := 0; i < 2*n; i++ {
        sanitizedI := i%n
        for len(st) != 0 && nums[sanitizedI] > nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            out[top] = nums[sanitizedI]
        }
        if i >= n {continue}
        st = append(st, sanitizedI)
    }
    return out
}

// processing ith element also works
// just run the loop 2 times
// func nextGreaterElements(nums []int) []int {
//     n := len(nums)
//     st := []int{} // vals    
//     out := make([]int, n)
//     for i := 0; i < n; i++ {out[i] = -1}

//     for i := (n-1)*2; i >= 0 ; i-- {
//         for len(st) != 0 && st[len(st)-1] <= nums[i%n] {
//             st = st[:len(st)-1]
//         }
//         if len(st) != 0 {
//             out[i%n] = st[len(st)-1]
//         }
//         st = append(st, nums[i%n])
//     }
//     return out
// }