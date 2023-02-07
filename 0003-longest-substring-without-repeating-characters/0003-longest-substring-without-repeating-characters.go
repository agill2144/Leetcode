/*
    classic sliding window - "find the length of the longest substring"
    
    sliding window type = Dynamic ( we need the longest one possible )
    
    approach:
    - have a left ptr ( slow ptr ) at idx 0 of string s
    - loop over string s ( with i pointer being the fast/right ptr)
    - We need to make sure that in each iteration, the characters between left/slow - right/fast ptrs are UNIQUE
    - How do we keep the unique-ness check?
        - Maintain a window map ( freqMap )
        - this map becomes responsible for maintaining the window state
        - each time we go over a character
        - if this character does not exist in windowMap, this means this char is uniq so far..
        - so simply add it to our map
        - get the current window size ( right/fast - left/slow + 1) 
            if currentSize > max { update max }
        - However if this character already exists in our windowMap, we need to slide our window such that this char is no longer part of our window
        - What do we need to update in our map when we slide our window?
            - decrement count in freqMap for each character that goes out of our window
            - We have 2 options
                - move left ptr 1 by 1 in our map data looks like {$key: $count} until the char to insert no longer exists in our map
                - or change the freqMap to idxMap {$char: idx}
                - and as soon as we run into a character that already exists in our map, move left/slow ptr to that charIdx+1
                - and update the idx of the char to be the new idx right/fast ptr is on 
    - return max
    
    time: o(n)
    space: o(n) 
*/
func lengthOfLongestSubstring(s string) int {
    left := 0
    maxSize := 0
    charIdxMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        idxLastSeenAt, ok := charIdxMap[char]
        if ok {
            if left <= idxLastSeenAt {
                left = idxLastSeenAt+1
            }
        }
        charIdxMap[s[i]] = i
        if i-left+1 > maxSize {
            maxSize = i-left+1
        }
    }
    return maxSize
}