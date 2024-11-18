func calculate(s string) int {
    total := 0
    curr := 0
    n := len(s)
    lastContr := 0
    var op byte = '+'
    for i := 0; i < n; i++ {
        if s[i] >= '0' && s[i] <= '9' {
            curr = curr * 10 + int(s[i]-'0')
        }
        if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' || i == n-1 {
            if op == '+' {
                total += curr
                lastContr = curr
            } else if op == '-' {
                total -= curr
                lastContr = -curr
            } else if op == '*'{
                total = total - lastContr + (lastContr * curr)
                lastContr = lastContr * curr
            } else if op == '/' {
                total = total - lastContr + (lastContr/curr)
                lastContr = lastContr / curr
            }
            curr = 0
            op = s[i]
        }
    }
    return total
}