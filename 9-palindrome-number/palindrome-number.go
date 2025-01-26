func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x != 0 && x % 10 == 0 {return false}
    rev := 0
    rightLimit := 214748364
    for x > rev {
        if rev > rightLimit || (rev == rightLimit && x % 10 > 7) {return false}
        rev = rev * 10 + (x % 10)
        x /= 10
    }
    return x == rev || rev / 10 == x
}