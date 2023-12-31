// we care about finding the longest possible substr between 2 identical chars
// therefore we need to keep track of each char and its idx position
// when we run into a char we have seen before, it means we have a possible substr
// check the size between the 2 positions ( idx from map and i and remove the outtermost 2 chars )
// and save this size if its bigger than our ans this far
// BECAUSE we want the longest possible size of a substr, we wont override idx positions for char
// that already exists in our idxMap, this will help us be greedy by keeping idx positions far away 
// time = o(s)
// space = o(1) ; we only have 26 keys at max (a to z)
func maxLengthBetweenEqualCharacters(s string) int {
    longest := math.MinInt64
    charIdx := map[byte]int{}
    for i := 0; i < len(s); i++ {
        char := s[i]
        idx, ok := charIdx[char]
        if ok {
            if (i-idx+1)-2 > longest {
                longest = (i-idx+1)-2
            }
        } else {
            charIdx[char] = i
        }
    }
    if longest == math.MinInt64 {return -1}
    return longest
}