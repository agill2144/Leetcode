// time: o(nlogn)
// space: o(1)
func reorderLogFiles(logs []string) []string {
    sort.SliceStable(logs, func(i, j int) bool{
        
        iLog := logs[i]
        jLog := logs[j]
        iLogSplit := strings.SplitN(iLog," ", 2) // "dig1 8 1 5 1" = ["dig1", "8 1 5 1"]
        jLogSplit := strings.SplitN(jLog," ", 2) // "g1 act car" = ["g1", "act car"]
        iLogIsDigitLog := isDigit(iLogSplit[1][0])
        jLogIsDigitLog := isDigit(jLogSplit[1][0])
        
        if !iLogIsDigitLog && !jLogIsDigitLog { // both are letter logs
            
            /*
                ["g2", "act car"]
                ["g1", "act car"]
            */
            if iLogSplit[1] == jLogSplit[1] { // if logs are same, sort by logID
                return iLogSplit[0] < jLogSplit[0]
            } else { // otherwise sort by log messages
                return iLogSplit[1] < jLogSplit[1]
            }
            
            
        } else if iLogIsDigitLog && !jLogIsDigitLog { 
            return false
        } else if !iLogIsDigitLog && jLogIsDigitLog {
            return true
        } else { 
            return false
        }        
    }) 
    return logs
}

func isDigit(n byte) bool {
    return n >= '0' && n <= '9'
}