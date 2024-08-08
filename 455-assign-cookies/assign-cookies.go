func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    cPtr := 0
    i := 0
    count := 0
    for i < len(g) && cPtr < len(s){
        if s[cPtr] >= g[i] {
            count++
            cPtr++
            i++
        } else {
            cPtr++
        }
    }
    return count
}