func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(start int, total int, lastContr int, path string)
    dfs = func(start int, total int, lastContr int, path string) {
        // base
        if start == len(num) {
            if total == target {
                out = append(out, path)
            }
            return
        }

        // logic
        for i := start; i < len(num); i++ {
            subStr := num[start:i+1]
            if len(subStr) > 1 && subStr[0] == '0' {continue}
            subStrInt, _ := strconv.Atoi(subStr)
            if len(path) == 0 {
                dfs(i+1, subStrInt, subStrInt, subStr)
            } else {
                dfs(i+1, total + subStrInt, subStrInt, path + "+" + subStr)
                dfs(i+1, total - subStrInt, -subStrInt, path + "-" + subStr)
                dfs(i+1, total - lastContr + lastContr * subStrInt, lastContr*subStrInt, path + "*" + subStr)
            }
        }
    }
    dfs(0,0,0,"")
    return out
}