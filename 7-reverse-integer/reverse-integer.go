func reverse(x int) int {
    // -2^31 = -2147483648
    // (2^31)-1 = 2147483647
    leftLimit := -214748364
    rightLimit := 214748364
    res := 0
    for x != 0 {
        lastDig := x%10
        if res < leftLimit || (res == leftLimit && lastDig < -8) {return 0}
        if res > rightLimit || (res == rightLimit && lastDig > 7) {return 0}
        res = res*10 + lastDig
        x /= 10
    }
    return res
}