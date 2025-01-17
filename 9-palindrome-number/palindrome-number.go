func isPalindrome(x int) bool {
    if x < 0 {return false}
    n := x
    rev := 0
    for n > rev {
        lastDig :=(n%10)
        if lastDig == 0 && rev == 0 {return false}
        rev = rev * 10 + lastDig
        n /= 10
    }
    return rev == n || rev/10 == n
}