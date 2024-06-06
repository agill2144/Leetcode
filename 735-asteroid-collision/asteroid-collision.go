func asteroidCollision(asteroids []int) []int {
    st := []int{}
    for i := 0; i < len(asteroids); i++ {
        curr := abs(asteroids[i])
        push := true
        for len(st) != 0 && st[len(st)-1] > 0 && asteroids[i] < 0 {
            top := st[len(st)-1]
            st = st[:len(st)-1]


            // when both are same size,
            // both are destroyed, nothing is pushed to st
            if top == curr {
                push = false
                break
            }

            // when top of st is bigger
            // curr is not pushed to st
            // top is pushed back
            if top > curr {
                st = append(st, top)
                push = false
                break
            }

            // when top is smaller than curr
            // continue removing
        }

        if push {st = append(st, asteroids[i])}
    }
    return st
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}