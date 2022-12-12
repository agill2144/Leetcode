/*
    approach : Two pointers
    - we can create 2 pointers p1 ( for word1 ) and p2 ( for word2 ), both with initial value of -1
    - these 2 pointers will hold the idx position of word1 and word2 when we find them
    - to find them, we will loop over wordsDict and in each iteration;
        - if we found word1, then set the p1 value to current idx
        - if we found word2, then set the p2 value to current idx
        - once both p1 and p2 are not -1, we have to positions to calculate distance from
        - calc dist and save it global min as needed
    - finally return min
    
    time: o(n)
    space: o(1)
*/

func shortestDistance(wordsDict []string, word1, word2 string) int {
    p1 := -1
    p2 := -1
    min := math.MaxInt64    
    for i := 0; i < len(wordsDict); i++ {
        if wordsDict[i] == word1 {
            p1 = i
        }
        if wordsDict[i] == word2 {
            p2 = i
        }
        if p1 != -1 && p2 != -1 {
            dist := abs(p1-p2)
            if dist < min {
                min = dist
            }
        }
        if min == 1 {
            break
        }
     }
    return min
}

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

// func shortestDistance(wordsDict []string, word1 string, word2 string) int {
//     idxMap := map[string][]int{}
//     for i := 0; i < len(wordsDict); i++ {
//         idxMap[wordsDict[i]] = append(idxMap[wordsDict[i]], i)
//     }
    
//     p1 := 0
//     p1List := idxMap[word1]
    
//     p2 := 0
//     p2List := idxMap[word2]
    
//     minDist := math.MaxInt64
//     for p1 < len(p1List) && p2 < len(p2List) {
//         p1Idx := p1List[p1]
//         p2Idx := p2List[p2]
//         dist := abs(p1Idx-p2Idx)
//         if dist < minDist {
//             minDist = dist
//         }
//         if p1Idx < p2Idx {
//             p1++
//         } else {
//             p2++
//         }
//     }
//     return minDist
// }

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}