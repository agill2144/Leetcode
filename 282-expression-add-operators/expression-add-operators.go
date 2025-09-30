func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(start int, total int, lastTerm int, expr string)
    dfs = func(start int, total int, lastTerm int, expr string) {
        // base
        if start == len(num) {
            if total == target {
                out = append(out, expr)
            }
            return
        }

        // logic
        for i := start; i < len(num); i++ {
            subStr := num[start:i+1]
            if len(subStr) > 1 && subStr[0] == '0' {return}
            val, _ := strconv.Atoi(subStr)
            if expr == "" {
                dfs(i+1, val, val, subStr)
            } else {
                dfs(i+1, total+val,+val,expr+"+"+subStr)
                dfs(i+1, total-val,-val, expr+"-"+subStr)
                multiRes := lastTerm * val                
                dfs(i+1, (total-lastTerm) + multiRes, lastTerm*val, expr+"*"+subStr)

            }
        }
    }
    dfs(0,0,0,"")
    return out
}
