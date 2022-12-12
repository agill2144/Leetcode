func characterReplacement(s string, k int) int {
    /*
       Return the length of the longest substring
       
       substring == sliding window
       longest == likely means we need to keep track to 2 points ( left and right )
                which sliding window does anyway..
                
                
        sliding window type:
        - Dynamic
                
        In our sliding window, we will keep track of a char freq map.
        Along that we will also keep track of the highest number of a repeated char ( the number )
        each time, we add a char to charFreqMap, we will check if the freq of this char from our map
        is greater than the max we have current have. Update it if its not.
        
        Then in each iteration, we have the number of chars seen so far ( r-l+1 )
        we also have the highest number a character has repeated in that window.
        Which means numCharsSeenSoFar - maxRepeats gives us the number of characters that would need to be replaced
        If those number of characters are <= k, save the window len.
        If those number of characters are >= k, shrink the window, this means we have too many other chars in current window that cannot be
        replaced with allocated k
        
        When shriking, we are simply updating our charFreqMap value
        
    */
    if len(s) <= 1 {
        return len(s)
    }
    
    charFreq := map[string]int{}
    left := 0
    maxRepeats := 0
    maxLen := 0
    for right := 0; right < len(s); right++{
        rightChar := string(s[right])
        charFreq[rightChar]++
        if val, _ := charFreq[rightChar]; val > maxRepeats {
            maxRepeats = val
        }
        
        numCharsSeenSoFar := right-left+1
        numReplacementsNeeded := numCharsSeenSoFar - maxRepeats
        
        if numReplacementsNeeded <= k {
            if right-left+1 > maxLen {
                maxLen = right-left+1
            }
        } else {
            leftChar := string(s[left])
            charFreq[leftChar]--
            if val, _ := charFreq[leftChar]; val == 0 {
                delete(charFreq, leftChar)
            }
            left++
        }
    }
    return maxLen
}