func asteroidCollision(asteroids []int) []int {
    st := []int{}    
    for i := 0; i < len(asteroids); i++ {
        curr := asteroids[i]
        pushCurr := true
        for len(st) != 0 && st[len(st)-1] > 0 && curr < 0 {
            top := st[len(st)-1]
            if abs(curr) >= top {
                st = st[:len(st)-1]
                if top == abs(curr) {pushCurr = false; break}
            } else {
                pushCurr = false
                break
            }
        }
        if pushCurr {
            st = append(st, asteroids[i])
        }
    }
    return st
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}