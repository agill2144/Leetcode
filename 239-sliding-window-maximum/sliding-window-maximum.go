func maxSlidingWindow(nums []int, k int) []int {
    dq := []int{} // idx
    out := []int{}
    left := 0
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        for len(dq) != 0 && curr > nums[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }
        dq = append(dq, i)
        if i-left+1 == k {
            out = append(out, nums[dq[0]])
            left++
            if len(dq) != 0 && dq[0] < left {
                dq = dq[1:]
            }
        }
    }
    return out
}