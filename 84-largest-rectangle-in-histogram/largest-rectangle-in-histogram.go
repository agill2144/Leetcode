func largestRectangleArea(heights []int) int {
    nsr := nsr(heights)
    nsl := nsl(heights)
    maxArea := 0
    for i := 0; i < len(heights); i++ {
        sr := nsr[i]
        sl := nsl[i]
        height := heights[i]
        width := sr-sl-1
        area := height*width
        maxArea = max(maxArea, area)
    }
    return maxArea
}

func nsr(nums []int) []int{
    n := len(nums)
    out := make([]int, n)
    st := []int{}
    for i := n-1; i >= 0; i-- {
        out[i] = n
        curr := nums[i]
        for len(st) != 0 && nums[st[len(st)-1]] >= curr {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            out[i] = st[len(st)-1]
        }
        st = append(st, i)
    }
    return out
}

func nsl(nums []int) []int{
    n := len(nums)
    out := make([]int, n)
    st := []int{}
    for i := 0; i < n ; i++ {
        out[i] = -1
        curr := nums[i]
        for len(st) != 0 && nums[st[len(st)-1]] >= curr {
            st = st[:len(st)-1]
        }
        if len(st) != 0 {
            out[i] = st[len(st)-1]
        }
        st = append(st, i)
    }
    return out
}