func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x != 0 && x % 10 == 0 {return false}
    tmp := x
    rev := 0
    for tmp != 0 {
        lastDig := tmp % 10
        rev = rev * 10 + lastDig
        tmp = tmp/10
    }
    return rev == x
}