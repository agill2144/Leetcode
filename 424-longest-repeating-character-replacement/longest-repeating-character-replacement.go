/*
    nested sliding window optimization:
    - can we use non shrinking window?
    - yes
    - when a window is valid, we can take it size and save it as needed
        - window is valid if
        - number of nonRepeating chars <= k
    - when is a window invalid?
        - when number of nonRepeating chars > k
        - this is when we will not shrink the size of our window
        - at the time same, we also WILL NOT SHRINK the maxFreq ( majority element )
        - shrinking maxFreq will never give us a better ans
        - a better ans is only possible when maxFreq of a char in a substr increases
        - therefore we are declaring maxFreq is part of non-shrinking window
            - just like window does not shrink, we only increase it 
            - similary, maxFreq will not decrease, it will only increase
    
    time = o(n)
    space = o(26) = o(1)
*/
func characterReplacement(s string, k int) int {
    freq := map[byte]int{}
    maxFreq := 0
    left := 0
    maxWin := 0
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
        maxFreq = max(maxFreq, freq[s[i]])
        nonRepeating := (i-left+1) - maxFreq
        if nonRepeating <= k {
            maxWin = max(maxWin, i-left+1)
        } else {
            freq[s[left]]--
            left++
        }
    }
    return maxWin

}

/*
    approach: sliding window
    - longest substring == sliding window
    - we have to find longest substr with same characters
    - but we can have k different chars
    - ABABA and k = 2
    - A repeats 3 times and B repeats 2 times
    - it makes sense to replace B's to A's so that we can have all A's
    - XGAAABAAX ; k = 3
        - GAAABAAX substr has 5 A's and 3 other chars
        - inorder to maximize same char substr, it makes sense to replace lower freq chars with maxFreq char
        - maxFreq in this one is A=5 and others are G=1; B=1; and X=1;
        - and we have 3 chars that we are allowed to replace
    - how to ensure we swap chars with lower freq in our window ? 
        - we maintain a freq map AND
        - we maintain a maxFreq of repeating char in a var
        - when we have a window to check;
        - we know its size=i-left+1
        - we also know how many times a char appears the most in this window = maxFreq
        - now, we need to find how many other chars are in this window that are not same as maxFreq char
        - nonRepeating := winSize - maxFreq
        - if we have a windowSize of 6 and 4 of them are all A's
        - how many chars are not A's ? 
            - 2; we get this from winSize - maxFreq
        - now if we k <= 2; this is a valid window
    - when our window becomes invalid;
        - shrink from left 1 by 1
        - reduce left char freq by 1
        - move left ptr forward
        - Now we may have a stale maxFreq number
        - therefore loop over freq map and find the next maxFreq val
    
    - time = o(2n) * 26
        = *26 is because of the extra loop we have to perform to find next maxFreq val
    - space = o(1);
*/
// func characterReplacement(s string, k int) int {
//     maxRepeating := 0
//     left := 0
//     maxWin := 0
//     freq := map[byte]int{}
//     for i := 0; i < len(s); i++ {
//         freq[s[i]]++
//         maxRepeating = max(maxRepeating, freq[s[i]])
//         nonRepeating := (i-left+1)-maxRepeating
//         for nonRepeating > k {
//             freq[s[left]]--
//             if freq[s[left]] == 0 {delete(freq, s[left])}
//             left++
//             maxRepeating := 0
//             for _, v := range freq {maxRepeating = max(maxRepeating, v)}
//             nonRepeating = (i-left+1)-maxRepeating
//         }
//         maxWin = max(maxWin, i-left+1)
//     }
//     return maxWin
// }