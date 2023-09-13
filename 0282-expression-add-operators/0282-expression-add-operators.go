func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(start, calc, contribution int, expr string)
    dfs = func(start, calc, contribution int, expr string) {
        // base
        if start == len(num) {
            if calc == target {
                out = append(out, expr)
            }
            return
        }
        
        // logic
        for i := start; i < len(num); i++ {
            currNumStr := string(num[start:i+1])
            if len(currNumStr) > 1 && currNumStr[0] == '0' {return}

            currNum, _ := strconv.Atoi(currNumStr)
            if len(expr) == 0 {
                dfs(i+1, currNum, currNum, currNumStr)    
            } else {
                // +
                dfs(i+1, calc+currNum, +currNum, expr + "+" + currNumStr)
                // -
                dfs(i+1, calc-currNum, -currNum, expr + "-" + currNumStr)
                // *
                dfs(i+1, (calc-contribution)+(currNum*contribution), currNum*contribution, expr + "*" + currNumStr)
            }
        }
    }
    
    dfs(0,0,0,"")
    return out
}