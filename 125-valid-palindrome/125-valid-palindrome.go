func isPalindrome(s string) bool {
    
    left, right := 0, len(s) - 1
    
    for left < right {
        l := rune(s[left])
        r := rune(s[right])

        // if char at left if not a letter or digit increment the left pointer
        if !unicode.IsLetter(l) && !unicode.IsDigit(l) {
            left++
        } else if !unicode.IsLetter(r) && !unicode.IsNumber(r) { // if char at right if not a letter or digit decrement the right pointer
            right--
        } else if unicode.ToLower(l) == unicode.ToLower(r) { // if both are same chars increment left pointer and decrement right pointer
            left++
            right--
        } else { // if none of the above conditions are satisified this is not a palindrome
            return false
        }
    }
    return true
    
    
}