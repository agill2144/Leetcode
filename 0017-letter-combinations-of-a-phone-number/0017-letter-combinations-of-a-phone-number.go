func letterCombinations(digits string) []string {
    if len(digits) == 0 {return nil}

    digMap := map[byte][]string{
        '2':[]string{"a","b", "c"},
        '3':[]string{"d","e", "f"},
        '4':[]string{"g","h", "i"},
        '5':[]string{"j","k", "l"},
        '6':[]string{"m","n", "o"},
        '7':[]string{"p","q", "r", "s"},
        '8':[]string{"t","u", "v"},
        '9':[]string{"w","x", "y", "z"},
    }

    out := []string{}
    var dfs func(ptr int, word string)
    dfs = func(ptr int, word string) {
        // base
        if ptr == len(digits) {
            out = append(out, word)
            return
        }
        
        // logic
        currDig := digits[ptr]
        chars := digMap[currDig]
        for j := 0; j < len(chars); j++ {
            word += chars[j]
            dfs(ptr+1, word)
            word = word[:len(word)-1]
        }
        
    }
    dfs(0, "")
    return out
}