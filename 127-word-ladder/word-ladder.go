func ladderLength(beginWord string, endWord string, wordList []string) int {
    set := map[string]bool{}
    foundEnd := false
    for i := 0; i < len(wordList); i++ {
        set[wordList[i]] = true
        if wordList[i] == endWord {foundEnd = true}
    }
    if !foundEnd {return 0}
    level := 1
    q := []string{beginWord}
    for len(q) != 0 {
        qSize := len(q)
        for qSize > 0 {
            dq := q[0]
            q = q[1:]
            if dq == endWord {return level}

            // explore all possible childs
            for i := 0; i < len(dq); i++  {
                for j := 0; j < 26; j++ {
                    child := dq[0:i] + string(j+'a') + dq[i+1:]
                    if set[child] {
                        q = append(q, child)
                        // set acts as a visited set as well
                        // once a child is added to queue
                        // that child is visited / no-longer usable
                        // therefore mark it visited 
                        // (set to false no longer allowed to use)
                        // because take a look at condition at line 22
                        // we only enqueue a child if its value is true in set
                        // by turning it false, we can no longer use it
                        // effectively marking it visited
                        set[child] = false
                    }
                }
            }
            qSize--
        }
        level++
    }
    return 0
}