/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    empMap := map[int]*Employee{}
    for i := 0; i < len(employees); i++ {
        empMap[employees[i].Id] = employees[i]
    }
    
    var dfs func(id int) int
    dfs = func(id int) int {
        // base
        
        // logic
        emp := empMap[id]
        importance := emp.Importance
        for i := 0; i < len(emp.Subordinates);  i++ {
            importance += dfs(emp.Subordinates[i])
        }
        return importance
    }
    return dfs(id)
}