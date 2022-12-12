// time: o(n)
// space: o(n)
func exclusiveTime(n int, logs []string) []int {
    result := make([]int, n)
    st := []int{}
    curr := 0
    prev := 0
    
    for _, log := range logs {
        logSplit := strings.Split(log,":")
        funcId,_ := strconv.Atoi(logSplit[0])
        logType := logSplit[1]
        t,_ := strconv.Atoi(logSplit[2])
        topIdx := len(st)-1
        curr = t
        if logType == "start" {
            if len(st) != 0 {
                result[st[topIdx]] += curr-prev
            }
            st = append(st, funcId)
            prev = curr
        } else {
            top := st[topIdx]
            st = st[:topIdx]
            result[top] += curr+1-prev
            prev = curr + 1
        }
    }
    return result
}