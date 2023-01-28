

/*
    - toss the entire dictonary in a trie
    - then split sentence into a list
    - Loop over each word in split array
    - Search for each character and form the replacement word
    - If we successfully found a replacement, then replace append the replacement to a our resulting string builder
    - Finally return the string from string builder
    

*/

// dictionary size = m, avg word len = k
// total words in sentence = n, avg word len = l
// total time = o(mk) + o(nl)
// space: o(mk) + o(nl)
func replaceWords(dictionary []string, sentence string) string {
    root := &trieNode{childrens: [26]*trieNode{}}
    // time : o(mk)
    for _, word := range dictionary {
        insert(root, word)
    }

    // time : o(n) * o(l) = o(nl)
    // o(l) where l is the avg len of each word in sentence and searching for a word in trie takes o(l) time
    // time = o(nl)
    out := new(strings.Builder)
    splitSen := strings.Split(sentence, " ")
    for idx, word := range splitSen {
        if idx != 0 {out.WriteString(" ")}
        found, replace := search(root, word)
        if found {
            splitSen[idx] = replace
        }
        out.WriteString(splitSen[idx])
    }
    
    
    return out.String()
}


type trieNode struct {
    isEnd bool
    childrens [26]*trieNode
}

// time: o(k) where k is len of input str
// space: o(k)
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

// time: o(k) where k is len of input str
// space: o(k) for string builder - worse case we do not find any smaller valid word
func search(root *trieNode, word string) (bool, string) {
    current := root
    out := new(strings.Builder)
    for _, char := range word {
        if current.childrens[char-'a'] == nil {return false, word}
        current = current.childrens[char-'a']
        out.WriteRune(char)
        if current.isEnd {
            return true, out.String()
        }
    }
    return false, word
}
