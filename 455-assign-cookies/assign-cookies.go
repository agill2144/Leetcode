func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    p1, p2 := 0, 0
    for p1 < len(g) && p2 < len(s) {
        if g[p1] <= s[p2] {
            p1++
            p2++
        } else {
            p2++
        }
    }
    return p1
}