/*
    approach: hash map
    - we need to search for word1 and word2 and find the closest distance between those 2 words
    - words could be repeated multiple times in the wordsDict
    - so for easy lookup of a word's idx, we will store word with all of its indicies in a map
    - then the map will look like: { $word: [$indicies] }
    - so we will loop over all elements in wordsDict and store them in the map
    - then grab the list of indicies for both word1 and word2
    - set two pointers on both of the list of indicies from the beginning
    - while both pointers are inbound,
        - calc abs dist between 2 indicies
        - save it in global min as needed
        - Now which pointer to move?
            - Since the indicies in both word1 and word2 list are going to be sorted
            - how are they already sorted?
            - because we went over the list from left to right when generating the idxMap
            - therefore already sorted.
            - moving which pointer is going to result into smaller diff?
            - if we move the pointer thats pointing to big value, the diff will only get bigger
            - if we move the pointer thats pointing to smaller value, there are chances of the diff to get smaller
            - therefore always move the smaller value pointer
    - finally return the min
    
    time: o(n)
    space: o(n) 
*/
type WordDistance struct {
    idxMap map[string][]int
}


func Constructor(wordsDict []string) WordDistance {
    idxMap := map[string][]int{}
    for idx, word := range wordsDict {
        idxMap[word] = append(idxMap[word], idx)
    }
    return WordDistance{idxMap: idxMap}
}


func (this *WordDistance) Shortest(word1 string, word2 string) int {
    
    w1Indices := this.idxMap[word1]
    w2Indices := this.idxMap[word2]
    
    w1 := 0
    w2 := 0
    minDist := math.MaxInt64
    for w1 < len(w1Indices)  && w2 < len(w2Indices) {
        
        w1Idx := w1Indices[w1]
        w2Idx := w2Indices[w2]
        
        diff := abs(w1Idx-w2Idx)
        if diff < minDist {
            minDist = diff
        }
        if minDist == 1 {
            break
        }
        
        if w1Idx < w2Idx {
            w1++
        } else {
            w2++
        }
    }
    return minDist
}


func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}

/**
 * Your WordDistance object will be instantiated and called as such:
 * obj := Constructor(wordsDict);
 * param_1 := obj.Shortest(word1,word2);
 */