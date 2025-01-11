func canConstruct(s string, k int) bool {
    if len(s) == k {return true}
    if len(s) < k {return false}
    freq := make([]int, 26)
    for i := 0; i < len(s); i++ {
        freq[int(s[i]-'a')]++
    }
    oddCount := 0
    for i := 0; i < len(freq); i++ {
        if freq[i] % 2 != 0 {oddCount++}
    }
    // oddCount are num of chars that need to stand on their own
    // if oddCount = 5, there would be 5 single char palindrome
    // remove k from oddCount, because we have to form k, 
    // if k was 3, and oddCount was 5, then have 2 oddCount left that are EXTRA!
    // therefore cannot construct
    // oddCount could have been smaller like 1, and k = 3
    // oddCount -= k ; 1 -= 3 = -2
    // which is fine, why ?
    // because the extra palindromes can still be formed 
    // using even characters, or by splitting characters redundantly into smaller palindromes.
    oddCount -= k
    return oddCount <= 0
}