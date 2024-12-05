func canChange(start string, target string) bool {
    s := 0
    t := 0
    for t < len(target) && s < len(start) {
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
    // if ptrs are still inbound, they have not failed yet
    // just process them all the way till the end and make sure they are all empty spaces
    for s < len(start) {if start[s] != '_'{return false}; s++}
    for t < len(target) {if target[t] != '_'{return false}; t++}
    return true
}