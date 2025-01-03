func asteroidCollision(ast []int) []int {
    n := len(ast)
    st := []int{}
    for i := 0; i < n; i++ {
        pushCurr := true
        for len(st) != 0 && st[len(st)-1] > 0 && ast[i] < 0 {
            if st[len(st)-1] == abs(ast[i]) {
                st = st[:len(st)-1]
                pushCurr = false
                break
            } else if st[len(st)-1] < abs(ast[i]) {
                st = st[:len(st)-1]
            } else if st[len(st)-1] > abs(ast[i]) {
                pushCurr = false
                break
            }
        }
        if pushCurr {
            st = append(st, ast[i])
        }
    }
    return st
}


func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}
