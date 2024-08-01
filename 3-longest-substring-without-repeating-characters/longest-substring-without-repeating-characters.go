/*
    brute force: 
    - make all possible substrings

    when we see substrings / subarrays == try sliding window

    sub optimal:
    - maintain a freq map of chars
    - if we are about to a add a char that already exists in our freq map
    - then move the left ptr until our current char does not exist in our window 
        and therefore does not exist in our freq map
    - once the above is satisfied, take the window size and save it as needed

    optimal:
    - when a char already exists in our freq map
    - we end up moving left ptr until the char no longer exists in freq map
    - so why not make left ptr take 1 big jump and move next to lastSeenIdx+1
    - now our freq map turns into a idx map
    - this way we can avoid the inner loop

    time = o(n)
    space = o(1) ; because at worse there are 256 chars in our map
*/
func lengthOfLongestSubstring(s string) int {
    idx := map[byte]int{}
    left := 0
    maxWindow := 0
    for i := 0; i < len(s); i++ {
        lastSeenAt, ok := idx[s[i]]
        if ok && left <= lastSeenAt {
            left = lastSeenAt+1
        }
        idx[s[i]] = i
        maxWindow = max(maxWindow, i-left+1)
    }
    return maxWindow
}