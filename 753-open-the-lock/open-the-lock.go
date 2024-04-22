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
                char := int(node[i]-'0')

                dec := char-1
                if char == 0 {dec = 9}

                
                child := node[:i] + fmt.Sprintf("%v",dec) + node[i+1:]
                _, ok := set[child]
                _, ok2 := visited[child]
                if !ok && !ok2 {
                    set[child] = struct{}{}
                    visited[child] = struct{}{}
                    q = append(q, child)
                }

                inc := char+1
                if char == 9 {inc = 0}
                child = node[:i] + fmt.Sprintf("%v",inc) + node[i+1:]
                _, ok = set[child]
                _, ok2 = visited[child]
                if !ok && !ok2 {
                    set[child] = struct{}{}
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