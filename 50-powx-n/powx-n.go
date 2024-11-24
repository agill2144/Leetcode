func myPow(x float64, n int) float64 {
    if n < 0 {
        x = 1/x
        n *= -1
    }
    var res float64 = 1.0
    // x^8 = (x^2)^4
    for n != 0 {
      if n % 2 != 0 {
        res *= x
        n--
      }
      if n <= 1 {break}
      x = x*x
      n = n/2
    }
    return res
}