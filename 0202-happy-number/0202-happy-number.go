func isHappy(n int) bool {
    totalMap := map[int]struct{}{}
    for true {
        total := 0
        for n != 0 {
            lastDig := n % 10
            total += lastDig * lastDig
            n /= 10
        }
        if total == 1 {return true}
        if _, ok := totalMap[total]; ok {return false}
        totalMap[total] = struct{}{}
        n = total
    }
    return false
}