// two ptrs, greedy
// if char matches, move both ptrs
// if char does not match, move sPtr until the char matches
//  handle the case of sPtr going out of bounds in this while loop
//  if after the while loop, sPtr is out of bounds, it means we have a char that does not exist in s
//      therefore we need the remaining of t string size as the min number of chars to be appended to s
//      therefore we can break
// once outside of the loop, there are 2 cases
//  1. all chars in t matched, meaning tPtr is at the end ( out-of-bounds )
//      in this case, we can exit and return 0 because we do not need to add more chars to t
//  2. some chars did not match, therefore the remaining size of the string is the min num of chars needed
// time = o(s) + o(t)
// space = o(1)
func appendCharacters(s string, t string) int {
    sPtr, tPtr := 0,0

    for sPtr < len(s) && tPtr < len(t) {
        if s[sPtr] == t[tPtr] {
            sPtr++
            tPtr++
        } else {
            for sPtr < len(s) && s[sPtr] != t[tPtr] {
                sPtr++
            }
            if sPtr == len(s) {break}
        }
    }
    if tPtr >= len(t) {return 0}
    return len(t) - tPtr
    
}