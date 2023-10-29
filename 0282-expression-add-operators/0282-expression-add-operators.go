func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(start, calc, contr int, expr string)
    dfs = func(start,calc,contr int, expr string) {
        // base
        if start == len(num) {
            if calc == target {
                out = append(out, expr)
            }
            return
        }
        
        // logic
        for i := start; i < len(num); i++ {
            subStr := num[start:i+1]
            if len(subStr) > 1 && subStr[0] == '0' {continue}
            subStrInt, _ := strconv.Atoi(subStr)
            if len(expr) == 0 {
                dfs(i+1, subStrInt, subStrInt, subStr)
            } else {
                dfs(i+1, calc+subStrInt, subStrInt, expr + "+" + subStr)
                dfs(i+1, calc-subStrInt, -subStrInt, expr + "-" + subStr)
                dfs(i+1, calc-contr+contr*subStrInt, contr*subStrInt, expr + "*" + subStr)
            }
        }
    }
    dfs(0,0,0,"")
    return out
}