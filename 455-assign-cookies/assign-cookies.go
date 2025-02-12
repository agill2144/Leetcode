/*
    approach: greedy
    - try to meet bare min requirement for each child
    time = o(g log g) + o(s log s) + o(max(s+g)) 
    space = o(g) + o(s) == if underlying sort is using merge sort 
*/
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