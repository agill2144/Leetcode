func canChange(start string, target string) bool {
    s := 0
    t := 0
    for t < len(target) || s < len(start) {
        for t < len(target) && target[t] == '_' {t++}
        for s < len(start) && start[s] == '_' {s++}
        if t == len(target) && s == len(start) {return true}
        if t == len(target) || s == len(start) {return false}
        if start[s] != target[t] {return false}
        if start[s] == 'L' {
            if t > s {return false}
        } else if start[s] == 'R' {
            if t < s {return false}
        }
        s++
        t++
    }
    return true
}