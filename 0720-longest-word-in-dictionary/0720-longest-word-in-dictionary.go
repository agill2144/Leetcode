/*
    approach: trie + DFS
    - Store all the words in a trie
    - Then loop thru all childrens of trie
        - 26 childrens because trie nodes contains characters ( for now )
        - and there are 26 characters ( lower case )
    - if a children is initialized and isEnd == true
    - add this children to our path ( the output string we are building to return )
    - Then recursively call the dfs on this child as the new root to explore other childs this child has initialized
    - In our base case, if len(path) > maxLen - then update the len and save the current word as a potential result
    - Finally once all the trie nodes are traversed, return the saved word
    
    - why dont we have to worry about the "smallest lexicographical order" in this ?
        - because we loop from a-z , therefore our path formation will already follow smallest lexicographical order
            
    
    time:
    - let n be the number of words
    - let k be the len of each word in words
    - o(nk) - to insert each word in a trie
    - 
*/

type trieNode struct{
    isEnd bool
    childrens [26]*trieNode
}

func insert(root *trieNode, word string) {
    current := root
    for _, char := range word {
        if current.childrens[char-'a'] == nil {
            current.childrens[char-'a'] = &trieNode{childrens: [26]*trieNode{}}
        }
        current = current.childrens[char-'a']
    }
    current.isEnd = true
}

func search(root *trieNode, word string) {
    
}

func longestWord(words []string) string {
    root := &trieNode{childrens: [26]*trieNode{}}
    for _, word := range words {
        insert(root, word)
    }
    var result string
    var backtrack func(r *trieNode, path string)
    backtrack = func(r *trieNode, path string) {
        // base
        if len(path) > len(result) {
            result = path
        }
        if r == nil {
            return
        }
        
        // logic
        for i := 0 ; i < 26; i++ {
            if r.childrens[i] != nil && r.childrens[i].isEnd{
                // action
                path += string('a'+i)
                // recurse
                backtrack(r.childrens[i], path)
                // backtrack
                path = path[:len(path)-1]
            }
        }
    }
    backtrack(root, "")
    return result
}