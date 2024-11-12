/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    adjList := map[int]*Employee{}
    for i := 0; i < len(employees); i++ {
        adjList[employees[i].Id] = employees[i]
    }
    total := 0
    var dfs func(node int)
    dfs = func(node int) {
        // base
        

        // logic
        total += adjList[node].Importance
        for _, sub := range adjList[node].Subordinates {
            dfs(sub)
        }
    }
    dfs(id)
    return total
}

/*
    {
        1: 1*Employee
        2: 2*Employee
        3: 3*Employee
    }
*/