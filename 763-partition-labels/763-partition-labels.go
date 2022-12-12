/*
    - pick a parition
    - ensure all characters in this partition are within this parition
    - if they are not, increase the parition size so that they are all within the parition size
    - save each partition size
    
    - now replace "partition" with a window
    - select a window to explore
    - ensure all characters in this window are within this window
    - if they are not, increase the window size so that they are all within the window size
    - save each window size once verified
    
    - question is how do we "select" a window/partition?
    - we will pick the first char and we will check where this characters last idx is
    - that will be our window to explore 
    - then we will go from idx 0 to the last idx of this window
    - and check each char by char and ensure that the last idx of this char is within the window ending idx
    
    - since we have to search for last idx of a character each time, lets build a charIdx map which keeps tracks of { $char: $endIdxOfThisChar }
    
    - The map helps us answer the question for each ith position in o(1) time
    - The map also helps us find the window size each time
*/
func partitionLabels(s string) []int {
    charIdx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        charIdx[s[i]]=i
    }
    
    result := []int{}
    start := -1
    windowEnd := -1
    
    for idx, char := range s {
        if start == -1 {
            start = idx
            windowEnd = charIdx[byte(char)]
        }
        currChar := char
        lastIdxOfThisChar := charIdx[byte(currChar)]
        if lastIdxOfThisChar > windowEnd {
            windowEnd = lastIdxOfThisChar
        }
        if idx == windowEnd {
            result = append(result, windowEnd-start+1)
            start = -1
        }
    }
    
    return result
}