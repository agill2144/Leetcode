func largestRectangleArea(heights []int) int {
   
    nsl := make([]int, len(heights))
    nsr := make([]int, len(heights))
    st := []int{}
    for i := 0; i < len(heights); i++ {
        h := heights[i]
        for len(st) != 0 && heights[st[len(st)-1]] >= h {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            nsl[i] = st[len(st)-1]+1
        }
        st = append(st, i)
    }

    st = []int{}
    for i := len(heights)-1; i >= 0; i-- {
        h := heights[i]
        for len(st) != 0 && heights[st[len(st)-1]] >= h {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            nsr[i] = st[len(st)-1]-1
        } else {
            nsr[i] = len(heights)-1
        }
        st = append(st, i)
    }
    ans := math.MinInt64
    for i := 0; i < len(heights); i++ {
        winSize := nsr[i] - nsl[i] + 1
        area := heights[i] * winSize
        ans = max(area, ans)
    }
    return ans
}