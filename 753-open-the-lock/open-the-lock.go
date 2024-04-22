func openLock(deadends []string, target string) int {
    start := "0000"
    set := map[string]struct{}{}
    for i := 0; i < len(deadends); i++ {
        if deadends[i] == start {return -1}
        set[deadends[i]] = struct{}{}
    }
    q := []string{start}
    level := 0
    // use the searchable deadends set as a visited map
    set[start] = struct{}{}
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            node := dq
            if node == target {
               return level
            }


            // find adj childs
            for i := 0; i < len(node); i++ {
                digit := int(node[i]-'0')
                turnLeft := digit-1
                if digit == 0 {turnLeft = 9}
                turnRight := digit+1
                if digit == 9 {turnRight = 0}

                c1 := node[:i] + fmt.Sprintf("%v", turnLeft) + node[i+1:]
                _, c1Visited := set[c1]
                if !c1Visited {
                    set[c1] = struct{}{}
                    q = append(q, c1)
                }
                c2 := node[:i] + fmt.Sprintf("%v", turnRight) + node[i+1:]                
                _, c2Visited := set[c2]
                if !c2Visited {
                    set[c2] = struct{}{}
                    q = append(q, c2)
                }
            }
            qSize--
        }
        level++
    }
    return -1
}