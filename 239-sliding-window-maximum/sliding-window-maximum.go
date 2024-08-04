func maxSlidingWindow(nums []int, k int) []int {
    dq := []int{} // idx
    left := 0
    out := []int{}
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        for len(dq) != 0 && curr > nums[dq[len(dq)-1]] {
            dq = dq[:len(dq)-1]
        }
        dq = append(dq, i)
        if i-left+1 == k {
            out = append(out, nums[dq[0]])
            if left == dq[0] {
                dq = dq[1:]
            }
            left++
        }
    }
    return out
}