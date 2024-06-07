// with trie
func replaceWords(dictionary []string, sentence string) string {
    root := &trieNode{childs: [26]*trieNode{}}
    for i := 0; i < len(dictionary); i++ {
        root.insert(dictionary[i])
    }
    out := new(strings.Builder)
    words := strings.Split(sentence, " ")
    for i := 0; i < len(words); i++ {
        word := words[i]
        found := false
        for j := 0; j < len(word); j++ {
            subStr := word[:j+1]
            if root.search(subStr) {
                found = true
                out.WriteString(subStr)
                break
            }
        }
        if !found {out.WriteString(word)}
        if i != len(words)-1 {out.WriteString(" ")}
    }
    return out.String()
}

type trieNode struct {
	isEnd  bool
	childs [26]*trieNode
}

func (r *trieNode) insert(word string) {
	curr := r
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if curr.childs[idx] == nil {
			curr.childs[idx] = &trieNode{childs: [26]*trieNode{}}
		}
		curr = curr.childs[idx]
	}
	curr.isEnd = true
}

func (r *trieNode) search(word string) bool {
	curr := r
	for i := 0; i < len(word); i++ {
		idx := int(word[i] - 'a')
		if curr.childs[idx] == nil {
			return false
		}
		curr = curr.childs[idx]
	}
	return curr.isEnd
}

// using hashset and without trie
// func replaceWords(dictionary []string, sentence string) string {
//     set := map[string]struct{}{}
//     for i := 0; i < len(dictionary); i++ {set[dictionary[i]] = struct{}{}}
//     out := new(strings.Builder)
//     words := strings.Split(sentence, " ")
//     for i := 0; i < len(words); i++ {
//         word := words[i]
//         found := false
//         for j := 0; j < len(word); j++ {
//             subStr := word[:j+1]
//             if _, ok := set[subStr]; ok {
//                 found = true
//                 out.WriteString(subStr)
//                 break
//             }
//         }
//         if !found {out.WriteString(word)}
//         if i != len(words)-1 {out.WriteString(" ")}
//     }
//     return out.String()
// }