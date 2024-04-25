/*
return number of substrings of len k 
with no dupe chars


fixed window

start with a map 
{$char:$freq}

if char to insert already exists in our map
then this means this a dupe char
and we need to move our left pointer forward

left pointer is going to move until the char to insert 
is no longer part of window
so instead of moving left pointer 1 by 1 and decrement freq count
of each char left points to - just move left to
lastSeenAtIdx+1 position

i.e store the idx in charMap instead of freq
if windowLen == k { count++ }



*/

func numKLenSubstrNoRepeats(s string, k int) int {
    
    if s == "" || k > len(s) {
		return 0
	}
	count := 0
	state := map[string]int{}
	left := 0
	// havefunonleetcode
	
	for right := 0; right < len(s); right++ {
		char := string(s[right])
		lastSeenAtIdx, ok := state[char]
		if ok {
			if left <= lastSeenAtIdx {
				left = lastSeenAtIdx + 1
			}
		}
		state[char] = right
		if right-left+1 == k {
			count++
			leftChar := string(s[left])
			if val := state[leftChar]; val <= left {
				delete(state, leftChar)
			}
			left++
		}
	}
	return count
    
}













