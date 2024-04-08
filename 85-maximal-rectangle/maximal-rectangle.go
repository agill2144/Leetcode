func maximalRectangle(matrix [][]byte) int {
    m := len(matrix)
    n := len(matrix[0])
    histogram := make([]int, n)
    maxArea := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '0' {
                histogram[j] = 0
            } else {
                histogram[j] += 1
            }
        }
        maxArea = max(maxArea, largestRectangleArea(histogram))
    }
    return maxArea

}


func largestRectangleArea(heights []int) int {
    nsl := nsl(heights)
    nsr := nsr(heights)
    ans := math.MinInt64
    for i := 0; i < len(heights); i++ {
        winSize := nsr[i] - nsl[i] + 1
        area := heights[i] * winSize
        ans = max(area, ans)
    }
    return ans
}

func nsr(arr []int) []int {
    out := make([]int, len(arr))
    st := []int{}
    for i := len(arr)-1; i >= 0; i-- {
        h := arr[i]
        for len(st) != 0 && arr[st[len(st)-1]] >= h {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            out[i] = st[len(st)-1]-1
        } else {
            out[i] = len(arr)-1
        }
        st = append(st, i)
    }
    return out
}

func nsl(arr []int) []int {
    out := make([]int, len(arr))
    st := []int{}
    for i := 0; i < len(arr); i++ {
        h := arr[i]
        for len(st) != 0 && arr[st[len(st)-1]] >= h {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            out[i] = st[len(st)-1]+1
        }
        st = append(st, i)
    }
    return out
}