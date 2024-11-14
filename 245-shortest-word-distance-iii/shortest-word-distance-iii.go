func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    w1 := -1; w2 := -1
    same := word1 == word2
    dist := len(wordsDict)
    for i := 0; i < len(wordsDict); i++ {
        if wordsDict[i] == word1 {
            if same {w2 = w1; w1 = i} else {w1 = i}
        } else if wordsDict[i] == word2 {
            w2 = i
        }
        if w1 != -1 && w2 != -1 {dist = min(dist, abs(w1-w2))}
    }
    return dist
}
/*
    approach: hashmap + two ptrs
    - store each word as key and their idx positions in a map[string][]int{}
    - then take 2 ptrs to loop over all idx positions for word1 and word2
    - if the words are same, then both ptrs are pointing at 0th idx;
        - therefore move w2 ptr 
    - now, do the classic two ptr
    - get dist between 2 idx positions ( values at w1 and w2 ptr )
    - save the calculated dist as needed
    - now, which ptr to move ?
    - if 2 words are same;
        - then we need to move both ptrs
    - otherwise
        - move the ptr with smaller idx value
        - moving the bigger idx value ptr will only make dist bigger
        - and we want smallest possible dist
        - hence move the smaller idx value ptr

    n = len(wordsDict)
    w1 = len(word1 idx positions)
    w2 = len(word2 idx positions)
    time = o(n) + o(w1 + w2)
    space = o(n); worst case, all words are unique
*/
// func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
//     idxs := map[string][]int{}
//     for idx, word := range wordsDict {
//         idxs[word] = append(idxs[word], idx)
//     }
//     same := word1 == word2
//     w1 := 0; w2 := 0
//     if same {w2++}
//     dist := math.MaxInt64
//     for w1 < len(idxs[word1]) && w2 < len(idxs[word2]) {
//         w1Idx := idxs[word1][w1]
//         w2Idx := idxs[word2][w2]
//         dist = min(dist, abs(w2Idx-w1Idx))
//         if same {w1++; w2++} else {
//             if w1Idx < w2Idx {w1++} else {w2++}
//         }
//     }
//     return dist
// }

func abs(x int) int {
    if x < 0 {return x * -1} 
    return x
}