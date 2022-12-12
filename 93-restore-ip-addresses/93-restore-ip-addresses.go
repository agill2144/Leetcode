func restoreIpAddresses(s string) []string {
    if len(s) == 0 {return nil}
    
    var result []string
    var backtrack func(start int, path []string)
    backtrack = func(start int, path []string) {
        // base
        if len(path) == 4 && start == len(s) {
            result = append(result, strings.Join(path,"."))
            return
        }
        if len(path) == 4 {return}
        
        // logic
        for i := start ; i < len(s); i++ {
            if s[start] == '0' && start != i {return}
            num := string(s[start:i+1])
            numInt, _ := strconv.Atoi(num)
            if numInt > 255 {continue}
            path = append(path, num)
            backtrack(i+1, path)
            path = path[:len(path)-1]
        }
    }
    backtrack(0,nil)
    return result
}