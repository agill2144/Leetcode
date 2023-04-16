/*
/*
    approach: topological sort - course schedule variation
    - This is follow up / related question of "verify alien dictonary"
        - in which we are given sorted words - sorted using the order
        - the order in which characters must be is also given
        - and we had to validate whether the list of words are sorted according the order of character provided.
        - In which we compare 2 words at a time from left to right
        - and if we run into a character mismatch - we make sure that word1 character comes before word2 character in the provided order
    
    - In this question, we have to return that order.
    - if we use the same strategy of comparing 2 words together, char by char, as soon as we see a mismatch
    - we can imply that word1 char MUST COME BEFORE word2 char
        - why can we imply the above?
        - because the question says "where the strings in words are sorted lexicographically by the rules of this new language"
    - So its safe to say that
        - word1Char is the parent and word2 char is the child
        - word1Char is the independant node ( does not depend on anyone ) and word2 char depends on word1 char ( its parent )
        - Essentially, word1Char is the parent node in an adjList and word2Char is the child that depends on word1Char
    - For each 2 word from the words dict, we will loop over all from left to right ( since they are apparentlt sorted in lexicographically )
    - and build our adjList
    - while we are building our adjList, we need to make sure that a child is not added to the same parent more than once 
        - since child here is just a char, this is likely to happen
        - so our adjList edges/childs will be a set of childs to ensure no duplicate child exists for the same parent
        - in a adjList/graph, the childs are the ones that have an incoming edge to them from their parent
        - Example: { "a": {"b","c"}  }  -- in a grap this would a -> b
                                                                  \> c
    - so while we are building out this adjList, we can also create our indegrees array
        - why is indegrees array needed?
        - indegrees array keep track of how many incoming edges a node has
            - i.e how many dependencies they have
            - in the above example; b and c depends on a, so therefore the number of dependencies b and c has is 1
            - which also means, that inorder to resolve b or c, its dependency must be resolved first
            - indegrees with 0 values - means that this idx node/char has no dependency
            - which means it can be proceessed without any issues - this is called the independent node 
                - i.e nodes with no incoming edges 
                - i.e nodes that do not depend on anyone
            - so we can process a first and once is proceessed - we need to update the dependants that were depending on a that a is now resolved
            - so therefore we can reduce a's dependants indegrees by 1
            - then b and c indegrees become 0 and therefore now they can be processed.
    - Now we know how and why we need a adjList and indegrees array
    - How are we going to do top sort?
    - Usually done using a BFS with a queue
        - enqueue independant nodes
        - update their dependants indegrees
        - enqueue dependants if their indegrees become 0.
    
    - Edge cases for this problem
        - the 2 words being compared for creating adjList and indegrees - may not be lexicographically sorted
            - so when we have 2 words and right before comparing character by character to build out the parent->{child} relationship
            - we will check whether 
                - word1 begins with word2
                - so we are treating ENTIRE word2 as a prefix
                - if word1 begins word2, and word1 len > word2 len - this is invalid 
                - this means that word1 should be placed after word2 for it to be truly lexicographically sorted.
                - in this case we will return nil for adjList and our indegrees array
*/
func alienOrder(words []string) string {
    adjList := map[byte][]byte{}
    
    // [ "a", "a" ]
    
    for i := 0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            adjList[words[i][j]] = []byte{}
        }
    }
    indegrees := make([]int, 26)
    for i := 0; i < len(words)-1; i++ {
        word1 := words[i]
        word2 := words[i+1]
        
        if strings.HasPrefix(word1, word2) && len(word2) < len(word1) {
            // invalid order
            return ""
        }
        
        for j := 0; j < len(word1) && j < len(word2); j++ {
            w1Char := word1[j]
            w2Char := word2[j]
            if w1Char != w2Char {
                adjList[w1Char] = append(adjList[w1Char], w2Char)
                indegrees[w2Char-'a']++
                break
            }
        }
    }
    q := []byte{}
    for key, _ := range adjList {
        if indegrees[key-'a'] == 0 {
            q = append(q, key)
        }
    }
    if len(q) == 0 {return ""}
    out := new(strings.Builder)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        out.WriteByte(dq)
        for _, nei := range adjList[dq] {
            indegrees[nei-'a']--
            if indegrees[nei-'a'] == 0 {
                q = append(q, nei)
            }
        }
    }
    // if there are still dependents not fully resolved, that means there was a cycle.
    for i := 0; i < 26; i++ {
        if indegrees[i] != 0 {return ""}
    }
    return out.String()
}