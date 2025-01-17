func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x % 10 == 0 && x != 0 {return false}
    rev := 0
    for x > rev {
        rev = rev * 10 + (x%10)
        x /= 10
    }
    return rev == x || rev/10 == x
}