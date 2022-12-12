/*
    WHEN YOU DONT HAVE TO USE ALL THE ELEMENTS TO MAKE A COMBINATION
    - FOR LOOP BASED RECURSION
    
    WHEN YOU HAVE TO USE ALL THE ELEMENTS TO MAKE A COMBINATION
    - MAKE DFS YOUR FOR LOOP WITH A START POINTER PICKING ELEMENTS FROM INPUT FORMING A PATH
    - ALTHOUGH FOR LOOP RECURSION WORKS TOO BUT HAVE TO BE MINDFUL 
    - WHETHER TO SELECT ONLY THE ith CHAR 
    - OR [0:i+1] CHAR (we did this in expression add operator - where we had multiple ints to be concatnated to evaluate n digit expression with other remaining numbers )
    
    Here we have to use each char for each digit, therefore making DFS act as my loop
    
*/  
func letterCombinations(digits string) []string {
    if len(digits) == 0 {return nil} 
    result := []string{}
    digMap := map[byte][]string{
        '1':[]string{},
        '2':[]string{"a","b", "c"},
        '3':[]string{"d","e", "f"},
        '4':[]string{"g","h", "i"},
        '5':[]string{"j","k", "l"},
        '6':[]string{"m","n", "o"},
        '7':[]string{"p","q", "r", "s"},
        '8':[]string{"t","u", "v"},
        '9':[]string{"w","x", "y", "z"},
    }
    var dfs func(path string, start int)
    dfs = func(path string, start int) {
        // base
        if start == len(digits) {
            result = append(result, path)
            return
        }
        
        
        // logic
        chars := digMap[digits[start]]
        for i := 0; i < len(chars); i++ {
            dfs(path+chars[i], start+1)
        }
    }
    dfs("", 0)
    return result
}