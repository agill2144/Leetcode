/*
    approach: reverse x (optimized)
    - do we need to reverse all digits?
    - if we had a string "1222221"
    - we would be comparing left char and right char ( using left and right ptrs )
    - until both ptrs have crosed each other
    - meaning we would traversing half of the input
    - can we do similar with our int input?
        - assume, we are not allowed to convert to string
    - how do we know we have reversed half of x ?
        - when rev version >= x
        - eg: x = 121, rev = 0
        - 1st iteration: x = 12, rev = 1
        - 2nd iteration: x = 1,  rev = 12
        - now rev version > x
        - what about uneven digits now?
        - x has 1 digit and rev has 2 digits
        - we have 2 cases;
            - 1. rev == x
            - 2. rev/10 == x
        - in this example; #2 applies
        - palindrome can have odd number of digits, because the middle one can be ignored

        - eg: 1221
        - 1st iteration: x = 122, rev = 1
        - 2nd iteration: x = 12,  rev = 12
        - rev >= x ? yes
        - x == rev ? yes
        - so for even cases, this works without needing to drop uneven digits from rev
*/
func isPalindrome(x int) bool {
    if x < 0 {return false}
    if x != 0 && x % 10 == 0 {return false}
    rev := 0
    for rev < x {
        rev = (rev * 10) + (x % 10)
        x /= 10
    }
    return x == rev || x == rev/10
}

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
// func isPalindrome(x int) bool {
//     // a negative number can never be a palidrome
//     // equality does not work
//     // eg: -121 -> rev() -> 121- 
//     if x < 0 {return false}
//     // when last digit is 0, reversing this num
//     // will never be equal to input number
//     // unless x itself is a zero
//     // eg: 10 -> rev() -> 01 -> 1 != 10
//     // eg: 120 -> rev() -> 021 -> 21 != 120
//     if x != 0 && x % 10 == 0 {return false}
//     rev := 0
//     tmp := x
//     for tmp != 0 {
//         lastDig := tmp % 10
//         rev = (rev*10)+ lastDig
//         tmp /= 10
//     }
//     return rev == x
// }