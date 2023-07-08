func lengthOfLongestSubstringKDistinct(s string, k int) int {
    
    
    /*
        return the length of the longest substring of s that contains at most k distinct characters.
        
        longest substr == sliding window
        
        window type:
        - Dynamic, since we are looking for the longest window possible
        
        So we need to find longest substr with at most k distinct chars -- i.e a substr with no more than k uniq chars
        
        So how do we keep track of number of characters seen so far?
        - map, a charFreq map
        
        What keeps increasing the window size?
        - as long as len(map) <= k
        
        When do we shrink our window?
        - if the len of our map > k
        - we will slide the window from left side and remove that left char freq count from our map
        - but that does not gurantee that our map will be <= k
        - we could have left ptr pointing to "a" and then another "a" after sliding -- still the map size wont be less than k
        - so then a while loop while len(map) > k 
        
        time: o(n)
        space: o(n)
    */
    
    if len(s) ==0 {
        return len(s)
    }
    
    // do not forget that hashing a string is actually o(k) time ... where k is the len of string 
    // but in this case we store char by char, so its o(1)
    // hashing funcs! dont forget how hashing functions are working for strings... 
    // we could use a ascii value too instead -- that works for char by char
    // wont work for words ( sum of all ascii values of all chars in words ) because collision - so prime product it instead... 
    charFreq := map[byte]int{}
    left := 0
    maxLen := 0
    for right := 0; right < len(s); right++ {
        rightChar := s[right]
        charFreq[rightChar]++
        if len(charFreq) <= k {
            if right-left+1 > maxLen {
                maxLen = right-left+1
            }
        } else {
            for len(charFreq) > k {
                leftChar := s[left]
                charFreq[leftChar]--
                if val, _ := charFreq[leftChar]; val == 0 {
                    delete(charFreq, leftChar)
                }
                left++
            }
        }
    }
    return maxLen
}