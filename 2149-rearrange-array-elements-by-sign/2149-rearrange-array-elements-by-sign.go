/*
    can we do it without extra space
    - when asked without extra space, the ans is either
        - use input array
        - or use output array 
    - in this case we can use output array
    - create an output arr of size n
    - create 2 ptrs
        1. positive ptr ; starting at idx 0
        2. negative ptr ; starting at idx 1
    - if nums[i] is positive, hand it to positive ptr, move this by 2
    - if nums[i] is negative, hand it to negative ptr, move this by 2
    - why are moving positive and negative by 2 ?
        - notice that we are placing with positive number first
        - i.e positive starts at idx 0, then the next positive will be at idx 1
        - because of alternating seq, idx 1 will contain a negative
        - therefore moving by 2
    
    time = o(n)
    space = o(1)
*/
func rearrangeArray(nums []int) []int {
    out := make([]int, len(nums))
    pPtr := 0 // positive ptr
    nPtr := 1 // negative ptr
    for i := 0; i < len(nums); i++{
        // collect all negatives
        if nums[i] < 0 {
            out[nPtr] = nums[i]
            nPtr+=2
        } else {
            out[pPtr] = nums[i]
            pPtr += 2
        }
    }
    return out
}   


/*
    approach: brute force
    - collect all positives in a list
    - collect all negatives in another list
    - use 2 ptrs to create alternating seq output array
    
    time = o(n)
    space = o(n)
*/
// func rearrangeArray(nums []int) []int {
//     p := []int{}
//     n := []int{}
//     for i := 0; i < len(nums); i++ {
//         if nums[i] > 0 {
//             p = append(p, nums[i])
//         } else {
//             n = append(n, nums[i])
//         }
//     }
//     out := []int{}
//     pPtr := 0
//     nPtr := 0
//     for pPtr < len(p) {
//         out = append(out, p[pPtr])
//         out = append(out, n[nPtr])
//         pPtr++
//         nPtr++
//     }
//     return out
// }