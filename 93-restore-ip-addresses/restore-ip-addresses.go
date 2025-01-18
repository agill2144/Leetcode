// numOfOptionsPerNode^totalNumOfOptions * extraWorkPerRecursion
// numOfOptionsPerNode = 3 ( substr size of 1,2,3)
// totalNumOfOptions = 4 ( we never exceed 4 chunks )
// extraWorkPerRecursion = o(4) always for strings.Join()
// tc = o(3^4) = o(1)
// sc = o(4) = o(1) ( recursion never exceeds depth of 4 )
func restoreIpAddresses(s string) []string {
    out := []string{}
    var dfs func(start int, path []string)
    dfs = func(start int, path []string) {
        // base
        if len(path) > 4 {return}
        if start == len(s) {
            if len(path) == 4 {
                out = append(out, strings.Join(path, "."))
            }
            return
        }

        // logic
        for i := start; i < len(s); i++ {
            subStr := s[start:i+1]
            if (len(subStr) > 1 && subStr[0] == '0') || len(subStr) > 3 {return}
            subStrInt, _ := strconv.Atoi(subStr)
            if subStrInt >= 0 && subStrInt <= 255 {
                path = append(path, subStr)
                dfs(i+1, path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, nil)
    return out
}