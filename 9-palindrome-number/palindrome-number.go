func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x % 10 == 0 && x != 0 {return false}
    n := x
    rev := 0
    for n > rev {
        rev = rev * 10 + (n%10)
        n /= 10
    }
    return rev == n || rev/10 == n
}