/*
    approach: brute force
    - collect all positives in a list
    - collect all negatives in another list
    - use 2 ptrs to create alternating seq output array
    
    time = o(n)
    space = o(n)
*/
func rearrangeArray(nums []int) []int {
    p := []int{}
    n := []int{}
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            p = append(p, nums[i])
        } else {
            n = append(n, nums[i])
        }
    }
    out := []int{}
    pPtr := 0
    nPtr := 0
    for pPtr < len(p) {
        out = append(out, p[pPtr])
        out = append(out, n[nPtr])
        pPtr++
        nPtr++
    }
    return out
}