func asteroidCollision(asteroids []int) []int {
    st := []int{}
    for i := 0; i < len(asteroids); i++ {
        curr := asteroids[i]
        if len(st) == 0 {
            st = append(st, curr)
            continue
        }
        
        pushCurr := true
        top := st[len(st)-1]
        for len(st) != 0 && top >= 0 && curr < 0 {
            
            if abs(curr) >= top{
                st = st[:len(st)-1]                
                if abs(curr) == top {pushCurr = false; break}
                if len(st) == 0 {break}
                top = st[len(st)-1]
            } else {
                pushCurr = false
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
    if x < 0 {return x * -1}
    return x
}