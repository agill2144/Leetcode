func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {return false}
    starts := []int{}
    for i := 0; i < len(goal); i++ {
        if s[i] == goal[i] {starts = append(starts, i)}
    }

    for k := 0; k < len(goal); k++ {
        sPtr := k
        g := len(goal)
        ok := true
        for i := 0; i < len(s); i++ {
            if s[i] != goal[sPtr%g] {ok = false; break}
            sPtr++
        }
        if ok {return true}
    }
    return false
}

// aacde cdeaa