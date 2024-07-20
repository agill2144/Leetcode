func maximumScore(nums []int, k int) int {
    n := len(nums)
    st := []int{}// idx
    maxScore := math.MinInt64
    for i := 0; i < len(nums); i++ {
        for len(st) != 0 && nums[i] < nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i-1
            nsl := 0
            if len(st) != 0 {
                nsl = st[len(st)-1]+1
            }
            if nsl <= k && k <= nsr {
                width := nsr-nsl+1
                score := nums[top] * width
                maxScore = max(maxScore, score)
            }
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsr := n-1
        nsl := 0
        if len(st) != 0 { nsl = st[len(st)-1]+1 }
        if nsl <= k && k <= nsr {
            width := nsr-nsl+1
            score := nums[top] * width
            maxScore = max(maxScore, score)
        }
    }
    return maxScore
}