func findContentChildren(greed []int, size []int) int {
    sort.Ints(greed)
    sort.Ints(size)

    g := 0
    s := 0
    for s < len(size) && g < len(greed) {
        cookieSize := size[s]
        greedFactor := greed[g]
        if greedFactor <= cookieSize {
            g++
            s++
        } else {
            s++
        }
    }
    return g

}