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
    var dfs func(digitPtr int, path string)
    dfs = func(digitPtr int, path string) {
        // base
        if digitPtr == len(digits) {
            out = append(out, path)
            return
        }

        // logic
        for i := 0; i < len(numberPad[digits[digitPtr]]); i++ {
            path += string(numberPad[digits[digitPtr]][i])
            dfs(digitPtr+1, path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, "")
    return out
}