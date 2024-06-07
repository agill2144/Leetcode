func largestRectangleArea(heights []int) int {
    st := []int{}
    maxArea := 0
    n := len(heights)
    for i := 0; i < n; i++ {
        curr := heights[i]
        for len(st) != 0 && curr < heights[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            nsr := i
            nsl := -1
            if len(st) != 0 {
                nsl = st[len(st)-1]
            }
            width := nsr-nsl-1
            height := heights[top]
            maxArea = max(maxArea, width*height)
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsr := n
        nsl := -1
        if len(st) != 0 {nsl = st[len(st)-1]}
        maxArea = max(maxArea, (nsr-nsl-1)*heights[top])
    }
    return maxArea
}