// greedy, start from the back and inverse the 2 ops
func brokenCalc(startValue int, target int) int {    
    count := 0
    for target > startValue {
        if target % 2 == 0 {
           target /= 2
        } else {
            target++
        }
        count++
    }
    return count + (startValue-target)
}