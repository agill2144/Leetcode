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
    for i := 0; i < len(employees); i++ {
        emp := employees[i]
        empRef[emp.Id] = emp
    }
    
    if _, ok := empRef[id]; !ok {return 0}
    
    sum := 0
    q := []*Employee{empRef[id]}
    for len(q) != 0 {
        dq := q[0]
        q = q[1:]
        sum += dq.Importance
        
        for _, empID := range dq.Subordinates {
            q = append(q, empRef[empID])
        }
    }
    return sum
}