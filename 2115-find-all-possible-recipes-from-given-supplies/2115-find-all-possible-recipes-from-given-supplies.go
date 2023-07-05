func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
 
    supplySet := newSet()
    recipeSet := newSet()
    for i := 0; i < len(recipes); i++ {recipeSet.add(recipes[i]); supplySet.add(recipes[i])}    
    for i := 0; i < len(supplies); i++ {supplySet.add(supplies[i])}
    
    indegrees := map[string]int{}
    adjList := map[string][]string{}
    for i := 0; i < len(recipes); i++ {
        allIng := ingredients[i]
        recipe := recipes[i]
        
        // // only if we have all ing
        // haveAll := true
        // for _, ing := range allIng {
        //     if !supplySet.contains(ing) {
        //         haveAll = false
        //         break
        //     }
        // }
        // if !haveAll {continue}
        indegrees[recipe] += len(allIng)
        
        for _, ing := range allIng {
            adjList[ing] = append(adjList[ing], recipe) 
        }
    }

    q := []string{}
    for _, supply := range supplies {
        q = append(q, supply)
    }
    
    out := []string{}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        
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