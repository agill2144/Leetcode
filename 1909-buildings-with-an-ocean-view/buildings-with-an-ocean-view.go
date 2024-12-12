func findBuildings(heights []int) []int {
    n := len(heights)
    out := []int{n-1}
    maxH := heights[n-1]
    for i := n-2; i >= 0; i-- {
        if heights[i] > maxH {
            out = append(out, i)
            maxH = heights[i]
        }
    }
    for i := 0; i < len(out)/2; i++ {
        out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
    }
    return out
}