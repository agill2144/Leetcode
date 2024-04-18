func minimumLength(s string) int {
    n := len(s)
    left := 0
    right := n-1

    for left < right {
        leftChar := s[left]
        rightChar := s[right]
        if leftChar == rightChar {
            // we can remove as many contagious occurences of the same char
            // move left as far in as possible
            // as long as new left == leftChar
            left++
            for left < right && s[left] == leftChar {
                left++
            }

            // move right as far back as possible
            // as long as new right == rightChar
            right--
            for left < right && s[right] == rightChar {
                right--
            }
        } else {
            break
        }
    }
    return right-left+1

}