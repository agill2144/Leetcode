func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    heights := make([]int, n)
    maxArea := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '0' {
                heights[j] = 0
            } else {
                heights[j] += 1
            }            
        }
        maxArea = max(maxArea, largestRectangleArea(heights))
    }
    return maxArea
}

// copy of: https://leetcode.com/problems/largest-rectangle-in-histogram/
func largestRectangleArea(heights []int) int {
    ans := math.MinInt64
    n := len(heights)
    st := []int{}
    for i := 0; i < n; i++ {

        // we are now only asking , am i(ith) your(top) nsr ? 
        for len(st) != 0 && heights[i] < heights[st[len(st)-1]] {
            height := heights[st[len(st)-1]]
            st = st[:len(st)-1]
            nsr := i // curr
            nsl := -1
            if len(st) != 0 {nsl = st[len(st)-1]}
            width := nsr-nsl-1
            area := height * width
            ans = max(ans, area)
        }
        st = append(st, i)
    }

    // if we never got a nsr ( for example, elements were in increasing order )
    // [1,2,3,4]
    // then for each ith element, we never got a true, that yes, ith element is top's nsr
    // therefore we will have all of these elments in our st
    // This means, each st's element nsr is len of the full heights array
    for len(st) != 0 {
        // remeber; we are processing the top element!
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
        ans = max(ans, area)
    } 
    return ans
}
