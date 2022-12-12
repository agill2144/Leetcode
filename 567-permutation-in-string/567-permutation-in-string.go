/*
    Sliding window problem
    - We need to find a window of size s1 in s2 that contains all characters of s1
    - They do not have to be in the same order as s1 ( it can be a permutation of s1 )
    - Sliding window type: Fixed
        - Len == len(s1)
    - Create a s1 map ( freqMap )
    - If we see a char while looping over s2 in s1 map, then we will reduce that char count by 1
    - Once we have reached a window of size len(s1) while looping over s2
    - We will check whether the value in freqMap is all 0's
    - If yes, it means our current window has all characters of s1
    - If no, slide the window forward
        - if the left char thats going out of the window is in s1 map, then increment its count to undo the decrement
    
    time: o(s1) + o(s2)
    space: o(1) // characters map will take a constant space of 26 characters
*/

func checkInclusion(s1 string, s2 string) bool {
    s1Map := map[string]int{}
    for _, char := range s1 {
        s1Map[string(char)]++
    }
    count := 0
    left := 0
    for right := 0; right < len(s2); right++ {
        rightChar := string(s2[right])
        if _, ok := s1Map[rightChar]; ok {
            s1Map[rightChar]--
            if val := s1Map[rightChar]; val == 0 {
                count++
            }
        }
        if right-left+1 == len(s1) {
            if count == len(s1Map) {
                return true
            }
            leftChar := string(s2[left])
            if val, ok := s1Map[leftChar]; ok {
                if val == 0 {count--}
                s1Map[leftChar]++
            }
            left++
        }
    }
    return false
}