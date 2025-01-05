func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(start int, total int, lastContr int, expr string)
    dfs = func(start, total, lastContr int, expr string) {
        // base
        if start == len(num){
            if total == target {out = append(out, expr)}
            return
        }

        // logic
        for i := start; i < len(num); i++ {
            subStr := num[start:i+1]
            if len(subStr) > 1 && subStr[0] == '0' {return}
            subStrInt, _ := strconv.Atoi(subStr)
            if expr == "" {
                dfs(i+1, subStrInt, subStrInt, subStr)
            } else {
                // +
                dfs(i+1, total+subStrInt, subStrInt, expr+"+"+subStr)
                // -
                dfs(i+1, total-subStrInt, -subStrInt, expr+"-"+subStr)
                // *
                dfs(i+1, total-lastContr+(lastContr*subStrInt), lastContr*subStrInt, expr+"*"+subStr)
            }
        }
    }
    dfs(0,0,0, "")
    return out
}