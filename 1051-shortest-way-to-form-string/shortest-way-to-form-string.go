func shortestWay(source string, target string) int {
    src := make([]bool, 26)
    for i := 0; i < len(source); i++ {src[int(source[i]-'a')] = true}
    s, t := 0, 0
    count := 0
    for t < len(target) {
        tChar := target[t]
        if !src[int(tChar-'a')] {return -1}
        sChar := source[s]
        if sChar == tChar {
            s++
            t++
        } else {
            s++
        }

        // either we have reached end of src, count++
        // or we have matched the entire target, count++
        if s == len(source) {
            s = 0
            count++
        } else if t == len(target) {
            count++
        }
    }
    return count
}