// numOfOptionsPerNode^totalNumberOfOptions
// numOfOptionsPerNode = 4 ( worst case we have digit 7 or 9 in hand, it has 4 chars, hence 4 options)
// totalNumberOfOptions = n = len(digits)
// tc = 4^n
// sc = n : recursion stack space will never exceed n size
func letterCombinations(digits string) []string {
    if len(digits) == 0 {return nil}
    out := []string{}
    numberPad := map[byte][]byte{
        '2': {'a', 'b', 'c'},
        '3': {'d', 'e', 'f'},
        '4': {'g', 'h', 'i'},
        '5': {'j', 'k', 'l'},
        '6': {'m', 'n', 'o'},
        '7': {'p', 'q', 'r', 's'},
        '8': {'t', 'u', 'v'},
        '9': {'w', 'x', 'y', 'z'},
    }
    var dfs func(ptr int, path string)
    dfs = func(ptr int, path string) {
        // base
        if ptr == len(digits) {out = append(out, path); return}

        // logic
        chars := numberPad[digits[ptr]]
        for i := 0; i < len(chars); i++ {
            dfs(ptr+1, path + string(chars[i]))
        }
    }
    dfs(0, "")
    return out
}