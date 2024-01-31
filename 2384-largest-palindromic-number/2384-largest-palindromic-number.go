/*
    hash + greedy:
    - we can use a digit to make palindrome if its freq is even 
        - 8's freq is 4; we can use all of the 8s
        - 8's freq is 5; we can use atleast 4 of them (i.e 5/2 )
        - the only time we cannot use a digit *right away* is if its freq is only 1
    - then lets build a freq map { $digit: $countTimes }
    - now we can loop over freq map
        - but not as naively as for k, v := range freqMap
        - THIS IS THE GREEDY PART
        - we want the "largest palindromic integer (in the form of a string)"
        - then lets be greedy and start from the highest possible number!
        - whats the highest possible digit = 9
        - and we run a loop from 9 to 0 ( inclusive )
        - then check if this digit is in our freqMap
        - if its freq is atleast >= 2 ; we can use this digit
    - now we want to build a string from outside -> in
        - i.e start placing palindromes on the outsides 
        - we can use a string builder that acts as the front
        - we can use a []string that acts as the back the string
            - []string for the back because we need the reverse order
            - so when building the final output string, we can loop from back of this and keep on appending to our final string
    - once we have process, we can STILL place 1 more digit in the middle of. $front and $back
    - since we want the highest possible integer palindrome, lets be greedy again!
        - range loop from 9 -> 0 ( inclusive )
        - check if this digit exists and has a freq of ATLEAST 1
        - once we find such digit, write to our final string builder ( acting as front of palindrome ) ; break and exit early from loop
    - finally write the back of the string in reverse order
    
    time = o(n) + o(9 * n) + o(n) = o(n)
    space = o(1) for freqMap only contains digits from 0 to 9 
          + o(n) for back of the string stored in []string
          = o(n)
*/
func largestPalindromic(num string) string {
    freq := map[int]int{}
    for i := 0; i < len(num); i++ {
        char := num[i]
        freq[int(char-'0')]++
    }
    out := new(strings.Builder)
    back := []string{}
    
    // form the front
    for i := 9; i >= 0; i-- {
        if freq[i] <= 1 {continue}
        count := freq[i] / 2
        
        freq[i] = freq[i] % 2
        for k := 0; k < count; k++ {
            if out.String() == "" && i == 0 {continue}
            out.WriteString(fmt.Sprintf("%v", i))
            // save the back chars
            back = append(back, fmt.Sprintf("%v", i))
        }
    }
    
    // form the middle
    for i := 9; i >= 0; i-- {
        if count, ok := freq[i]; ok && count >= 1 {
            out.WriteString(fmt.Sprintf("%v", i))
            break
        }
    }
    
    for i := len(back)-1; i >= 0; i-- {
        out.WriteString(back[i])
    }
    if out.String() == "" {out.WriteString("0")}
    return out.String()
    
}
