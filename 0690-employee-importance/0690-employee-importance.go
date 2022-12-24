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
    sum := 0
    var dfs func(e *Employee)
    dfs = func(e *Employee) {
        // base
        if e == nil {return}
        
        // logic
        sum += e.Importance
        subordinates := e.Subordinates
        for _, sub := range subordinates {
            dfs(empRef[sub])
        }
    }
    dfs(start)
    return sum
}