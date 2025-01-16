func isPalindrome(x int) bool {
    if x < 0 {return false}
    n := x
    rev := 0
    for n != 0 {
        rev = (rev*10)+ n%10
        n /= 10
    }
    return rev == x
}