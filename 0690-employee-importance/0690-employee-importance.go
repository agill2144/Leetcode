/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    if employees == nil || len(employees) == 0 {return 0}

    empRef := map[int]*Employee{}
    var start *Employee
    for i := 0; i < len(employees); i++ {
        emp := employees[i]
        empRef[emp.Id] = emp
        if emp.Id == id {
            start = emp
        }
    }
    var dfs func(e *Employee) int
    dfs = func(e *Employee) int {
        // base
        if e == nil {return 0}
        
        // logic
        sum := e.Importance
        subordinates := e.Subordinates
        for _, sub := range subordinates {
            sum += dfs(empRef[sub])
        }
        return sum
    }
    return dfs(start)
}