func maxSlidingWindow(nums []int, k int) []int {
    dq := []int{}
    left := 0
    out := []int{}
    for i := 0; i < len(nums); i++ {
        for len(dq) != 0 && nums[i] > nums[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }
        dq = append(dq, i)
        if i-left+1 == k {
            out = append(out, nums[dq[0]])
            left++
            if dq[0] < left {dq = dq[1:]}
        }
    }
    return out
}