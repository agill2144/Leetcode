func customSortString(order string, s string) string {
    idx := map[string]int{}
    for i := 0; i < len(order); i++ {
        idx[string(order[i])] = i
    }
    chars := strings.Split(s, "")
    sort.Slice(chars, func(i, j int)bool {
        iidx, iok := idx[chars[i]]
        jidx, jok := idx[chars[j]]
        if iok && jok {
            return iidx < jidx
        } else if iok {
            return true
        } else if jok {
            return false
        }
        return true
    })
    return strings.Join(chars, "")
}