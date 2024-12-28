func asteroidCollision(asteroid []int) []int {
    st := []int{}
    for i := 0; i < len(asteroid); i++ {
        pushCurr := true
        for len(st) != 0 && asteroid[i] < 0 && st[len(st)-1] > 0 {
            if abs(asteroid[i]) <= st[len(st)-1] {
                if abs(asteroid[i]) == st[len(st)-1] {st = st[:len(st)-1]}
                pushCurr = false
                break
            }
            st = st[:len(st)-1]                
        }
        if pushCurr {
            st = append(st, asteroid[i])
        }
    }
    return st
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}