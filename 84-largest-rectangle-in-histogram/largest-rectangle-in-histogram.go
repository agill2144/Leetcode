func largestRectangleArea(heights []int) int {
    n := len(heights)
    maxArea := 0
    st := []int{} // idx
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
            maxArea = max(maxArea, height*width)
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        nsr := n
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsl := -1
        if len(st) != 0 {
            nsl = st[len(st)-1]
        }
        width := nsr-nsl-1
        height := heights[top]
        maxArea = max(maxArea, height*width)
    }
    return maxArea
}