func reverseWords(s []byte)  {
    // reverse entire thing
    for i := 0; i < len(s)/2; i++ {
        s[i], s[len(s)-1-i] = s[len(s)-1-i],s[i]
    }
    s = append(s,  ' ') 
    // reverse each word
    slow := 0
    for i := 0; i < len(s); i++{
        if s[i] == ' ' {
            reverse(s, slow, i-1)
            slow = i+1
        }
    }
}

func reverse(s []byte, left, right int) {
    for left < right {
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}