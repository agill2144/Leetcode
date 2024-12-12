func myPow(x float64, n int) float64 {
    if n == 0 {return 1.0}
    if n < 0 {
        x = 1/x
        n *= -1
    }
    var extraTerms float64 = 1.0
    for n > 1 {
        if n % 2 == 0 {
            x *= x
            n /= 2
        } else {
            extraTerms *= x            
            n--
        }
    }
    return x * extraTerms
}