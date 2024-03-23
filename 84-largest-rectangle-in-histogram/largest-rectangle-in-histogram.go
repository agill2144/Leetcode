func largestRectangleArea(heights []int) int {
    type stackNode struct {
        idx int
        val int
    }
    nsl := make([]int, len(heights))
    nsr := make([]int, len(heights))
    st := []stackNode{}
    for i := 0; i < len(heights); i++ {
        h := heights[i]
        for len(st) != 0 && st[len(st)-1].val >= h {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            nsl[i] = st[len(st)-1].idx+1
        }
        st = append(st, stackNode{i,h})
    }

    st = []stackNode{}
    for i := len(heights)-1; i >= 0; i-- {
        h := heights[i]
        for len(st) != 0 && st[len(st)-1].val >= h {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            nsr[i] = st[len(st)-1].idx-1
        } else {
            nsr[i] = len(heights)-1
        }
        st = append(st, stackNode{i, h})
    }
    ans := math.MinInt64
    for i := 0; i < len(heights); i++ {
        winSize := nsr[i] - nsl[i] + 1
        area := heights[i] * winSize
        ans = max(area, ans)
    }
    return ans
}