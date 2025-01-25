func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    i, j := 0, 0
    for i < len(g) && j < len(s) {
        if s[j] >= g[i] {i++; j++; continue}
        j++
    }
    return i
}