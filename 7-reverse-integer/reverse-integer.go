func reverse(x int) int {
    // left limit = -2147483648
    left := -214748364
    // right limit = 2147483647
    right := 214748364
    rev := 0
    for x != 0 {
        lastDig := x % 10
        if rev < left || (rev == left && lastDig > 8) {return 0}
        if rev > right || (rev == right && lastDig > 7) {return 0}
        rev = rev * 10 + lastDig
        x /= 10
    }
    return rev
}