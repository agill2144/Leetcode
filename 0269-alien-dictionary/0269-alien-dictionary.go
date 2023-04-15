func alienOrder(words []string) string {
    result := new(strings.Builder)
    
    adjList, indegrees := buildGraph(words)
    q := []byte{}
    for char, _ := range adjList {
        if indegrees[char-'a'] == 0 {
            q = append(q, char)
        }
    }
    // if we have no independent node(s) to start with , return empty string
    if len(q) == 0 {
        return ""
    }
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        result.WriteByte(dq)
        childSet := adjList[dq]
        for child, _ := range childSet.items {
            indegrees[child-'a']--
            if indegrees[child-'a'] == 0 {
                q = append(q, child)
            }
        }
    }
    outStr := result.String()
    // if we have not used all characters from adjList ( keys ), that means there
    // are still nodes in indegrees to process that could not be processed
    // meaning there is no valid top sort
    // i.e return ""
    if len(outStr) != len(adjList) {
        return ""
    }
    return outStr
}

func buildGraph(words []string) (map[byte]*set, []int) {
    out := map[byte]*set{}
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            out[words[i][j]] = newSet()
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        
        /*
        func HasPrefix(s, prefix string) bool
        HasPrefix tests whether the string s begins with prefix.
        */
        if strings.HasPrefix(string(word1), string(word2)) && len(word2) < len(word1) {
            return map[byte]*set{}, nil
        }
        
        for j := 0 ; j < len(word1) && j < len(word2); j++ {
            w1Char := word1[j]
            w2Char := word2[j]
            if w1Char != w2Char {
                // w1 is the parent/independent
                // w2 is the child ( i.e it has an incoming edge )
                if !out[w1Char].contains(w2Char){
                    indegrees[w2Char-'a']++
                    out[w1Char].add(w2Char)
                }
                break
            } 
        }
    }
    return out, indegrees   
}

type set struct{ items map[byte]struct{} }
func newSet() *set { return &set{items: map[byte]struct{}{}} }
func (s *set) add(x byte){ s.items[x] = struct{}{} }
func (s *set) remove(x byte) { delete(s.items, x)}
func (s *set) contains(x byte) bool { _, ok := s.items[x]; return ok}