/*
    approach: DFS
    - For each alphabet character we have 2 things to do
        - lower case
        - upper case
    - Why not a for loop based recursion with a start pointer ?
        - that works too, however this will cause only the ith to be added to our path in each recursion
        - which leads to not using all characters of the string
        - and we have to use all characters of the string
        - if we use ith idx from for loop to pick out characters from input string, NOT ALL CHARACTERS WILL BE PICKED
        - example: a1b2
        - When the path has "a" and the for loop ith idx is at idx 2 ( so at char "b")
        - path + char[i] will contain "ab" -- you see how we do not have "1"
            - then we will reurse with dfs(ab) and dfs(aB) -- sure we can have our base case to only add to result if the len path == len string
            - but these are extra dfs calls being made that can be avoided.
    - Instead we will make our dfs act as a loop
    - In which our dfs will take a ith idx value to take action on
        - action: add the character to our path
    - if the character is a digit, we do not have 2 choices, so simply add it to our path as it and recurse to the next idx
    - otherwise
        - 2 dfs recursive calls will be made
        - first with lower case character appended to our path
        - then with upper case character appended to our path
    - Since we cannot reuse the characters, we will make our i pointer in each recursive call move by +1
    - At some point this ith position will be out of bounds - since the i pointer is being used to pick out characters from s string 1 by 1 
    - When this happens, we have all the characters from input string in our path, save the result, and return the call back to parent func
    - Do we need to backtrack anything?
        - no, not in golang, since strings are pass by values
        - strings will be copied and original strings from previous calls will not be modified in memory because of pass by value
    
    
    
    WHEN YOU DONT HAVE TO USE ALL THE ELEMENTS TO MAKE A COMBINATION
    - FOR LOOP BASED RECURSION
    
    WHEN YOU HAVE TO USE ALL THE ELEMENTS TO MAKE A COMBINATION
    - MAKE DFS YOUR FOR LOOP WITH A START POINTER PICKING ELEMENTS FROM INPUT FORMING A PATH
    - ALTHOUGH FOR LOOP RECURSION WORKS TOO BUT HAVE TO BE MINDFUL 
    - WHETHER TO SELECT ONLY THE ith CHAR 
    - OR [0:i+1] CHAR (we did this in expression add operator - where we had multiple ints to be concatnated to evaluate n digit expression with other remaining numbers )
    
    
    time: 2^n
    space:  2^n
*/
func letterCasePermutation(s string) []string {
    result := []string{}
    var dfs func(start int, path string)
    dfs = func(start int, path string) {
        // base
        if start == len(s) {
            result = append(result, path)
            return
        }
        
        // logic
        char := s[start]
        strChar := string(char)
        if char >= '0' && char <= '9' {
            dfs(start+1, path+strChar)
        } else {
            dfs(start+1, path+strings.ToUpper(strChar))
            dfs(start+1, path+strings.ToLower(strChar))
        }
    }
    dfs(0, "")
    return result
}