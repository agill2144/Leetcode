func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    heights := make([]int, n)
    maxArea := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                heights[j] += 1
            } else {
                heights[j] = 0
            }
        }
        maxArea = max(maxArea, largestRectangleArea(heights))
    }
    return maxArea
}

func largestRectangleArea(heights []int) int {
    st := []int{} // indices
    maxArea := 0
    for i := 0; i < len(heights); i++ {
        curr := heights[i]
        for len(st) != 0 && curr < heights[st[len(st)-1]] {
            // process top
            nsr := i
            top := st[len(st)-1]
            st = st[:len(st)-1]
            // if there are no elements inside the stack, it means
            // this bar can be expanded all the way to the left
            // therefore setting initial nsl idx = -1
            nsl := -1
            if len(st) != 0 {
                nsl = st[len(st)-1]
            }
            width := nsr-nsl-1
            area := heights[top] * width
            maxArea = max(area, maxArea)
        }
        st = append(st, i)
    }

    for len(st) != 0 {
        nsr := len(heights)
        top := st[len(st)-1]
        st = st[:len(st)-1]
        nsl := -1
        if len(st) != 0 {
            nsl = st[len(st)-1]
        }
        width := nsr-nsl-1
        area := heights[top] * width
        maxArea = max(area, maxArea)        
    }
    return maxArea
}
