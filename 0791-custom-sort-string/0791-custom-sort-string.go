/*
    we want to sort s based on a custom "order"
    therefore we should collect all chars from order string that also exist in "s"
    because order matters, chars in $order must be placed first
    
    appraoch: using a hashset
    - store chars of s in a set
    - loop thru the order string to collect chars that should be placed first
    - if the order[i] does exist in s set, take it and add it to our output
        - delete this char from set
        - so that we know for sure its used
    - then loop over s string one last time, to collect remaining chars that did not exist in $order
    - there is a possibility that not all chars of $s exist in $order,
    - therefore we must also preserve the order of chars in $s that do not exist in $order
    
    s = len(s)
    t = len(order)
    
    time = o(s) + o(t) + o(s)
    space = o(s)
*/
func customSortString(order string, s string) string {
    
    // make the chars of $s searchable 
    freqMap := map[byte]int{}
    for i := 0; i < len(s); i++ {
        freqMap[s[i]]++
    }
    
    out := new(strings.Builder)
    // collect the chars from order that also exist in $s
    for i := 0; i < len(order); i++ {
        char := order[i]
        count := freqMap[char]
        for count > 0 {
            out.WriteByte(char)
            count--
        }
        delete(freqMap, char)
    }
    
    // collect the remaining left over chars from $s that did not exist in $order
    for i := 0; i < len(s); i++ {
        count := freqMap[s[i]]
        if count > 0 {
            out.WriteByte(s[i])
        }
    }
    return out.String()
}
