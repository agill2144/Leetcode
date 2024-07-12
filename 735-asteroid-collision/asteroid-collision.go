func asteroidCollision(asteroids []int) []int {
    st := []int{}
    n := len(asteroids)
    for i := 0; i < n; i++ {
        curr := asteroids[i]
        pushCurr := true
        for len(st) != 0 && st[len(st)-1] > 0 && curr < 0 {
            top := st[len(st)-1]
            st = st[:len(st)-1]
            // top == abs(curr)
            if top == abs(curr) {
                pushCurr = false
                break
            } else if top > abs(curr) {
                pushCurr = false
                st = append(st, top)
                break
            }
        }
        if pushCurr {
            st = append(st, curr)
        }
    }
    return st
}

func abs(x int) int {
    if x < 0 {return x *-1 }
    return x
}