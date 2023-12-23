func isPathCrossing(path string) bool {
    set := map[[2]int]struct{}{}
    set[[2]int{0,0}] = struct{}{}
    x := 0
    y := 0
    for i := 0; i < len(path); i++ {
        if path[i] == 'N' {
            y++
        } else if path[i] == 'S' {
            y--
        } else if path[i] == 'E' {
            x++
        } else if path[i] == 'W' {
            x--
        }
        cord := [2]int{x, y}
        if _, ok := set[cord]; ok {
            return true
        }
        set[cord] = struct{}{}
    }
    return false
}