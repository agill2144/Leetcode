func canMakeSubsequence(str1 string, str2 string) bool {
    s2 := 0
    for i := 0; i < len(str1) && s2 < len(str2); i++ {
        s1Char := int(str1[i]-'a')
        s2Char := int(str2[s2]-'a')
        if s1Char == s2Char {s2++; continue}
        if (s1Char+1)%26 == s2Char {s2++}
    }
    return s2 == len(str2)
}

func abs(x int) int {
    if x < 0 {return x*-1}
    return x
}