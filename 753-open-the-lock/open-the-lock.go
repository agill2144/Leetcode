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

            // node = 4 char string
            // each char is a wheel on the locker
            // each wheel can be rotated once forward or backward
            for i := 0; i < len(node); i++ {
                digit := int(node[i] - '0')
                
                // Generate the forward and backward combinations
                for _, d := range []int{-1, 1} {
                    nextDigit := (digit + d + 10) % 10
                    child := node[:i] + strconv.Itoa(nextDigit) + node[i+1:]
                    if _, ok := set[child]; ok { continue }
                    if _, ok := visited[child]; ok { continue }
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