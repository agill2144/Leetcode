func maximumScore(nums []int, k int) int {
    st := []int{}
    maxScore := 0
    for i := 0; i < len(nums); i++ {
        curr := nums[i]
        for len(st) != 0 && curr < nums[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i-1
            nsl := 0
            if len(st) != 0 {nsl = st[len(st)-1]+1}
            width := nsr-nsl+1
            score := nums[top] * width
            if nsl <= k && k <= nsr {
                maxScore = max(maxScore, score)
            }
        }
        st = append(st, i)
    }
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]        
        nsr := len(nums)-1
        nsl := 0
        if len(st) != 0 {nsl = st[len(st)-1]+1}
        width := nsr-nsl+1
        score := nums[top] * width
        if nsl <= k && k <= nsr {
            maxScore = max(maxScore, score)
        }

    }
    return maxScore
}