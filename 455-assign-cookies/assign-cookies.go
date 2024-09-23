/*
    approach: greedy
    - try to meet bare min requirement for each child
    time = o(g log g) + o(s log s) + o(max(s+g)) 
    space = o(g) + o(s) == if underlying sort is using merge sort 
*/
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