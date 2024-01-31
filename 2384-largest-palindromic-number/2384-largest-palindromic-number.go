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
        - and we run a loop from 9 to 0 (inclusive )
        - then check if this digit is in our freqMap
        - if its freq is >= 2
    - now we want to build a string from outside -> in
        - i.e start placing palindromes on the out sides 
        - we can use a string builder that acts as the front
        - we can use a []string that acts as the back the string
            - []string for the back because we need the reverse order
            - so when building the final output string, we can loop from back of this and keep on appending to our final string
    - once we have process
*/
func largestPalindromic(num string) string {
    freq := map[int]int{}
    for i := 0; i < len(num); i++ {
        char := num[i]
        freq[int(char-'0')]++
    }
    out := new(strings.Builder)
    var middle string
    back := []string{}
    for i := 9; i >= 0; i-- {
        if freq[i] <= 1 {continue}
        count := freq[i] / 2
        
        freq[i] = freq[i] % 2
        for k := 0; k < count; k++ {
            if out.String() == "" && i == 0 {continue}
            out.WriteString(fmt.Sprintf("%v", i))
            back = append(back, fmt.Sprintf("%v", i))
        }
    }
    for i := 9; i >= 0; i-- {
        if count, ok := freq[i]; ok && count >= 1 {
            middle = fmt.Sprintf("%v", i)
            break
        }
    }
    out.WriteString(middle)
    for i := len(back)-1; i >= 0; i-- {
        out.WriteString(back[i])
    }
    if out.String() == "" {out.WriteString("0")}
    return out.String()
    
}
