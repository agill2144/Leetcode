func findBuildings(heights []int) []int {
    n := len(heights)
    maxOnRight := heights[n-1]
    out := []int{n-1}
    for i := n-2; i >= 0; i-- {
        if heights[i] > maxOnRight {
            out = append(out, i)
            maxOnRight = heights[i]
        }
    }
    for i := 0; i < len(out)/2; i++ {
        out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
    }
    return out
}