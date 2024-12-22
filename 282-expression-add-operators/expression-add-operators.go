func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(start int, res int, prevContr int, exp string)
    dfs = func(start int, res int, prevContr int, exp string) {
        // base
        if start == len(num) {
            if res == target {
                out = append(out, exp)
            }
            return
        }

        // logic
        for i := start; i < len(num); i++ {
            numStr := num[start:i+1]
            // EDGE CASE!
            // WHENEVER CONSTRUCTING NUMBERS, CHECK FOR LEADING ZERO CASES!
            // substr could have more than 1 digit
            // but substr starts with 0
            // which is not a valid number
            if len(numStr) > 1 && numStr[0] == '0' {continue}
            n, _ := strconv.Atoi(numStr)
            if len(exp) == 0 {
                dfs(i+1, n, n, numStr)
            } else {
                dfs(i+1, res+n, n, exp+"+"+numStr)
                dfs(i+1, res-n, -n, exp+"-"+numStr)
                dfs(i+1, (res-prevContr)+(prevContr*n), prevContr*n, exp+"*"+numStr)
            }
        }
    }
    dfs(0,0,0,"")
    return out
}