func canConstruct(ransomNote string, magazine string) bool {
    bucket := make([]int, 26)
    for i := 0; i < len(magazine); i++ {
        bucket[int(magazine[i]-'a')]++
    }
    for i := 0; i < len(ransomNote); i++ {
        bucket[int(ransomNote[i]-'a')]--
        if bucket[int(ransomNote[i]-'a')] < 0 {return false}
    }
    return true
}