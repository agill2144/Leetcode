func openLock(deadends []string, target string) int {
    start := "0000"
    set := map[string]struct{}{}
    for i := 0; i < len(deadends); i++ {
        if deadends[i] == start {return -1}
        set[deadends[i]] = struct{}{}
    }
    q := []string{start}
    level := 0
    visited := map[string]struct{}{}
    visited[start] = struct{}{}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq
            if node == target {
               return level
            }

            for i := 0; i < len(node); i++ {
                // for each wheel, we can do +1(forward) or -1(backward)
                digit := int(node[i]-'0')
                choices := []int{1,-1}
                for _, ch := range choices {
                    newDigit := (digit + ch + 10) % 10
                    child := node[:i] + fmt.Sprintf("%v", newDigit) + node[i+1:]
                    if _, ok := set[child]; ok {continue}
                    if _, ok := visited[child]; ok {continue}
                    visited[child] = struct{}{}
                    q = append(q, child)
                }
            }
            qSize--
        }
        level++
    }
    return -1
}