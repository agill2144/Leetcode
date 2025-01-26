func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x != 0 && x % 10 == 0 {return false}
    rev := 0
    for x > rev {
        rev = rev * 10 + (x % 10)
        x /= 10
    }
    return x == rev || rev / 10 == x
}