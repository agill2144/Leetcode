func largestRectangleArea(heights []int) int {
    st := []int{}
    maxArea := 0
    n := len(heights)
    // top of st is processed
    // if ith element is the nsr of top

    for i := 0; i < n; i++ {
        curr := heights[i]
        for len(st) != 0 && curr < heights[st[len(st)-1]] {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            height := heights[top]
            nsr := i
            nsl := -1
            if len(st) != 0 {
                nsl = st[len(st)-1]
            }
            width := nsr-nsl-1
            area := height*width
            maxArea = max(maxArea, area)
        }
        st = append(st, i)
    }
    // edge case when processing top of st;
    // we have items left in the stack, when would this happen ?
    // when none of the ith values were a NSR of top
    // this means, nsr of top element does not exist, it means we can use the entire len
    for len(st) != 0 {
        top := st[len(st)-1]
        st = st[:len(st)-1]
        height := heights[top]
        nsr := n
        nsl := -1
        if len(st) != 0 {
            nsl = st[len(st)-1]
        }
        width := nsr-nsl-1
        area := height*width
        maxArea = max(maxArea, area)
    }
    return maxArea
}
