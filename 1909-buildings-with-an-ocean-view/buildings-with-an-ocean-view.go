func findBuildings(heights []int) []int {
    st := []int{} // buildings on right side with greater height so far
    n := len(heights)
    out := []int{}
    out = append(out, n-1)
    st = append(st, heights[n-1])
    for i := n-2; i >= 0; i-- {
        if heights[i] > st[len(st)-1] {
            st = st[:len(st)-1]
            out = append(out, i)
            st = append(st, heights[i])
        }
    }
    for i := 0; i < len(out)/2; i++ {out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]}
    return out
}