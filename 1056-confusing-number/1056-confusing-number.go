func confusingNumber(n int) bool {
    flipMap := map[int]int{0:0,1:1,6:9,8:8,9:6}
    orig := n
    flip := 0
    for n != 0 {
        last := n%10
        n = n/10
        flipValOfLast , ok := flipMap[last]
        if !ok {
            return false
        }
        flip = flip*10+flipValOfLast
    }
    return orig != flip
}