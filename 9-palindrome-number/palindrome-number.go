// tc = o(k) = k is num of digits in x
// sc = o(1)
func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x != 0 && x % 10 == 0 {return false}
    rev := 0
    for x > rev {
        lastDig := x % 10
        rev = rev * 10 + lastDig
        x /= 10
    }
    return rev == x || rev/10 == x
}

/**

    x = 1
    rev = 12 

    x == rev || x == rev / 10

*/