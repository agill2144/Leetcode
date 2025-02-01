func findBuildings(heights []int) []int {
    n := len(heights)
    out := []int{}
    maxH := -1
    for i := n-1; i >= 0; i-- {
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