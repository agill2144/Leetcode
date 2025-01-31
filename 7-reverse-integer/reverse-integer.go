func reverse(x int) int {
    res := 0
    // hard limit -2147483648
    leftLimit := -214748364
    // hard limit = 2147483647
    rightLimit := 214748364

    for x != 0 {
        lastDig := x % 10
        x /= 10
        if res < leftLimit || (res == leftLimit && lastDig > 8) {return 0}
        if res > rightLimit || (res == rightLimit && lastDig > 7) {return 0}
        res = res * 10 + lastDig        
    }
    return res
}