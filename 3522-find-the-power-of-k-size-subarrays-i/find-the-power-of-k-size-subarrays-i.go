func resultsArray(nums []int, k int) []int {
    dq := []int{0} // idx
    n := len(nums)
    left := 0
    out := []int{}
    for i := 0; i < n; i++ {
        if len(dq) == 0 {
            dq = append(dq, i)
        } else if nums[i] == nums[dq[len(dq)-1]]+1 {
            dq = append(dq, i)
        } else {
            dq = []int{i}
        }

        if i-left+1 == k {
            if len(dq) == k {
                out = append(out, nums[dq[len(dq)-1]])
            } else {
                out = append(out, -1)
            }
            for len(dq) > 0 && dq[0] <= left {dq = dq[1:]}
            left++
        }
    }
    return out
}