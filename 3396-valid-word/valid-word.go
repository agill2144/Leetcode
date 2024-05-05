func isValid(word string) bool {
    if len(word) < 3 {return false}
    vowelCheck := false
    consonantCheck := false
    for i := 0; i < len(word); i++ {
        char := word[i]
        if isVowel(char) {
            vowelCheck = true
        } else if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            consonantCheck = true
        }
        if (char >= '0' && char <= '9') || ( char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
            continue
        }

        return false
    }
    return vowelCheck && consonantCheck
}

func isVowel(char byte) bool {
    return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' 
}