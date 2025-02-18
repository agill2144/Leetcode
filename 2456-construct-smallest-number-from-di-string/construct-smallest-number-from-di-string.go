func smallestNumber(pattern string) string {
    res := ""
    st := []int{}
    count := 1
    for i := 0; i <= len(pattern); i++ {
        st = append(st, count)
        count++
        if i == len(pattern) || pattern[i] == 'I' {
            for len(st) != 0 {
                res += fmt.Sprintf("%v",st[len(st)-1])
                st = st[:len(st)-1]
            }
        }
    }
    return res
}