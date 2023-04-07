func trap(height []int) int {
    slow := 0
    running := 0
    total := 0
    for i := 1; i < len(height); i++ {
        slowH := height[slow]
        currH := height[i]
        if currH < slowH {
            running += slowH-currH
        } else {
            total += running
            running = 0
            slow = i
        }
    }
    running = 0
    pivot := slow
    slow = len(height)-1
    for i := len(height)-1; i >= pivot; i-- {
        if i == slow {continue}
        currH := height[i]
        slowH := height[slow]
        if currH < slowH {
            running += slowH-currH
        } else {
            total += running
            slow = i
            running = 0
        }
    }
    return total
}