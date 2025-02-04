func exclusiveTime(n int, logs []string) []int {
    out := make([]int, n)
    st := [][]int{}
    
    for i := 0; i < len(logs); i++ {
        log := strings.Split(logs[i], ":")
        funcId, _ := strconv.Atoi(log[0])
        logType := log[1]
        logTime, _ := strconv.Atoi(log[2])
        
        if logType == "start" {
            // If stack is not empty, update the time of the previous function
            if len(st) > 0 {
                prevFuncId := st[len(st)-1][0]
                prevStartTime := st[len(st)-1][1]
                out[prevFuncId] += logTime - prevStartTime
            }
            // Push the current function onto the stack
            st = append(st, []int{funcId, logTime})
        } else { // "end" type
            // Pop the top function from the stack
            topLog := st[len(st)-1]
            st = st[:len(st)-1]
            
            // Calculate and add the exclusive time for this function
            topFuncId := topLog[0]
            topStartTime := topLog[1]
            out[topFuncId] += logTime - topStartTime + 1
            
            // If there are still functions on the stack, adjust their start time
            if len(st) > 0 {
                st[len(st)-1][1] = logTime + 1
            }
        }
    }
    
    return out
}