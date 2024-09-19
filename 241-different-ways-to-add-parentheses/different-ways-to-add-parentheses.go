func diffWaysToCompute(expression string) []int {
    var dfs func(expr string) []int
    dfs = func(expr string) []int {
        // base
        if len(expr) <= 2 {
            if expr[0] >= '0' && expr[0] <= '9' {
                exprInt,  _ := strconv.Atoi(expr)
                return []int{exprInt}
            }
        }

        // logic
        res := []int{}
        for i := 0 ; i < len(expr); i++ {
            if expr[i] == '+' || expr[i] == '-' || expr[i] == '*' {
                left := dfs(expr[0:i])
                right := dfs(expr[i+1:])

                for j := 0; j < len(left); j++ {
                    for k := 0; k < len(right); k++ {
                        if expr[i] == '+' {
                            res = append(res, left[j]+right[k])
                        } else if expr[i] == '-' {
                            res = append(res, left[j]-right[k])
                        } else if expr[i] == '*' {
                            res = append(res, left[j]*right[k])
                        }
                    }
                }
            }
        }

        return res
    }
    return dfs(expression)
}