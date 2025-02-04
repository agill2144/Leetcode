func myPow(x float64, n int) float64 {
    // x^n  = x^n/2 * x^n/2
    if n < 0 {
        n *= -1
        x = 1/x
    }
    var res float64 = 1.0
    for n != 0 {
        if n % 2 == 0 {
            x *= x
            n /= 2
        } else {
            res *= x
            n--
        }
    }
    return res
}