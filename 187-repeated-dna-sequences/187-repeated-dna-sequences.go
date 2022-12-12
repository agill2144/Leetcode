/*
    find all 10 len substrings that are repeated more than once

    sliding window ( without the left pointer ... that wont be needed )
    fixed window size -- 10 len
    How do we keep track of window state to be compared later?
    - We will keep a window state and we want to do quickly search
    - therefore a map[string]bool{} -- bool as a flag to dedupe results

    The approach is that we will keep looping over the string and forming a temp string.
    as soon as the len of temp string == 10,
    we will check in our window state whether we have seen this or not.
    If we have not, add it. and reset the tmp to "" and move one.
    If we have, then check if the value is false, if it is, save this to output array and update the value of this key to be true (as in successfully saved so no one else re-uses this value again and adds a dupe entry in our output array )


    time: o(n)
    space: o(n)
*/
func findRepeatedDnaSequences(s string) []string {
    state := map[string]bool{}
    str := ""
    out := []string{}
    for right := 0; right < len(s); right++ {
        str += string(s[right])
        if len(str) == 10 {
            used, ok := state[str]
            if !ok {
                state[str] = false
            } else if !used {
                out = append(out, str)
                state[str] = true
            }
            str = str[1:]
        }
    }
    return out
}