func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    g1 := 0
    s1 := 0
    for g1 < len(g) && s1 < len(s) {
        size := s[s1]
        greed := g[g1]
        if size >= greed {
            s1++
            g1++
        } else if size < greed {
            s1++
        }
    }
    if g1 == len(g) {return len(g)}
    return g1
}