func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    p1 := 0
    p2 := 0
    for p1 < len(g) && p2 < len(s) {
        greed := g[p1]
        size := s[p2]
        if greed <= size {
            p1++
            p2++
        } else {
            p2++
        }
    }
    return p1
}