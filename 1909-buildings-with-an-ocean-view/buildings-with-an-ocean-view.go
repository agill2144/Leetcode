func findBuildings(heights []int) []int {
    st := []int{} // idxs
    for i := 0; i < len(heights); i++ {
        for len(st) != 0 && heights[i] >= heights[st[len(st)-1]] {
            st = st[:len(st)-1]
        }
        st = append(st, i)
    }
    return st
}