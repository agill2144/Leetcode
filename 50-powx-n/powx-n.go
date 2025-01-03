func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n *= -1
    }
    // x^n = (x^n/2) * (x^n/2) 

    var res float64 = 1.0
    for n != 0 {
        if n % 2 == 0 {
            x *= x
            n /= 2
        } else {
            n--
            res *= x
        }
    }
    return res
}