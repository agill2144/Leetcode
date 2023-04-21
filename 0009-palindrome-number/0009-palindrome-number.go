func isPalindrome(x int) bool {
    if x < 0 {return false}
    n := x
    revr := 0
    for x != 0 {
        lastDig := x % 10
        x /= 10
        revr = revr * 10 + lastDig
    }
    // fmt.Println(revr)
    return revr == n
}