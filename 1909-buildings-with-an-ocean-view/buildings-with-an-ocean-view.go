func findBuildings(heights []int) []int {
    n := len(heights)
    st := []int{}    
    for i := 0; i < n; i++ {
        for len(st) != 0 && heights[i] >= heights[st[len(st)-1]] {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return st
}