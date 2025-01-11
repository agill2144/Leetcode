func checkValidString(s string) bool {
    asteriks := []int{} // idxs of asteriks
    st := []int{} // idxs of open parans
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            st = append(st, i)
        } else if s[i] == ')' {
            if len(st) > 0 {
                st = st[:len(st)-1]
                continue
            }
            if len(asteriks) > 0 {
                asteriks = asteriks[:len(asteriks)-1]
                continue                
            }
            return false
        } else if s[i] == '*' {
            asteriks = append(asteriks, i)
        }
    }
    for len(st) != 0 && len(asteriks) != 0 {
        if asteriks[len(asteriks)-1] > st[len(st)-1] {
            asteriks = asteriks[:len(asteriks)-1]
            st = st[:len(st)-1]
            continue
        }
        break
    }
    return len(st) == 0
}