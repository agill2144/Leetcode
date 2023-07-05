func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
 
    /*
        v = num of recipes
        
        s = num of supplies
        k = avg num of ingredients per recipe
    */
    
    // time and space = o(r)
    recipeSet := newSet()
    for i := 0; i < len(recipes); i++ {recipeSet.add(recipes[i])}
    
    
    indegrees := map[string]int{} // space = o(r)
    adjList := map[string][]string{} // space = o(k+r)

    // time = o(k+r)
    for i := 0; i < len(recipes); i++ {
        allIng := ingredients[i]
        recipe := recipes[i]
        indegrees[recipe] += len(allIng)
        
        // recipe depends on its ingredients
        // therefore each ing is independent 
        // therefore ingredient -> recipe
        // adjList = {independent : [dependents]}
        // adjList = {$ing : [$recipe]}
        
        for _, ing := range allIng {
            adjList[ing] = append(adjList[ing], recipe) 
        }
    }

    q := []string{}
    // why start from supplies and not with indegrees with 0 ?
    // ingredients list could also have other recipes 
    // therefore ingredients list could contain supplies + other recipe names
    // supplies are on their own standing, no edges, no indegrees, they do not have any recipes within them
    // therefore starting with supplies
    
    // time = o(s)
    // space in queue initially = o(s)
    for _, supply := range supplies {
        q = append(q, supply)
    }
    
    out := []string{}
    
    // time = o(i+v)
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        
        // is this a recipe? if yes, add it, as its successfully created
        if recipeSet.contains(dq) {
            out = append(out, dq)
        }
        for _, nei := range adjList[dq] {
            if _, ok := indegrees[nei]; ok {
                indegrees[nei]--
                if indegrees[nei] == 0 {
                    q = append(q, nei)
                }
            }
        }
    }
    return out
}


type set struct {
    items map[string]struct{}
}
func newSet() *set {
    return &set{items: map[string]struct{}{}}
}
func (s *set) add(x string) {
    s.items[x] = struct{}{}
}
func (s *set) remove(x string) {
    delete(s.items, x)
}
func (s *set) contains(x string) bool {
    _, ok := s.items[x]
    return ok
}