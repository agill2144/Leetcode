/*
    approach: reverse x
    - for a number to be palindrome
    - it must have the same value when read from left->right and from right->left
    - eg: 11, 22, 121, 939, 2, 3, are valid palindromes
    - eg: 10, 12, 93, 123, -121 are not valid palindromes
    - to reverse a int value
    - get the last digit by performing $x % 10
    - then update x with left over digits = $x / 10
    - eg: 123
    - lastDigit = 123 % 10 = 3 is the remainder and also our last digit
    - updated x = x / 10 = 12 is the quotient
    - we do this until x becomes 0

    k = num of digits
    tc = o(k)
    sc = o(1)
*/
func isPalindrome(x int) bool {
    // a negative number can never be a palidrome
    // equality does not work
    // eg: -121 -> rev() -> 121- 
    if x < 0 {return false}
    // when last digit is 0, reversing this num
    // will never be equal to input number
    // unless x itself is a zero
    // eg: 10 -> rev() -> 01 -> 1 != 10
    // eg: 120 -> rev() -> 021 -> 21 != 120
    if x != 0 && x % 10 == 0 {return false}
    rev := 0
    tmp := x
    for tmp != 0 {
        lastDig := tmp % 10
        rev = (rev*10)+ lastDig
        tmp /= 10
    }
    return rev == x
}