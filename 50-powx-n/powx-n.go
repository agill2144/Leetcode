// x^n = (x*x)^n/2
// PEMDAS 
func myPow(x float64, n int) float64 {
    var ans float64 = 1.0
    if n < 0 {n*=-1; x = 1/x}
    for n != 0 {
        if n % 2 == 0 {
            x *= x
            n /= 2
        } else {
            ans *= x
            n--
        }
    }
    return ans
}