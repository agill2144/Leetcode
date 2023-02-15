func addOperators(num string, target int) []string {
    out := []string{}
    var dfs func(path string, start , calc, tail int)
    dfs = func(path string, start, calc, tail int) {
        // base
        if start == len(num) {
            if calc == target {
                out = append(out, path)
            }
            return
        }
        // logic
        for i := start; i < len(num); i++ {
            // if the sub number we are forming, starts with 0, this is an invalid number
            curr := num[start:i+1]
            if len(curr) > 1 && curr[0] == '0' {continue}
            currNum,_ := strconv.Atoi(curr)

            if start == 0 {
                dfs(string(curr),i+1, currNum, currNum)
            } else {
                // 3 options : +, -, x
                dfs(path+"+"+string(curr), i+1, calc+currNum, +currNum)
                dfs(path+"-"+string(curr), i+1, calc-currNum, -currNum)
                dfs(path+"*"+string(curr), i+1, calc-tail+(tail*currNum), tail*currNum)
            }
        }
    }
    dfs("", 0, 0, 0)
    return out
}