package solution690

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) (total int) {
	m := map[int]*Employee{}

	for _, employee := range employees {
		m[employee.Id] = employee
	}

	queue := []int{id}

	for len(queue) > 0 {
		employee := m[queue[0]]
		queue = queue[1:]
		total += employee.Importance
		queue = append(queue, employee.Subordinates...)
	}

	return
}
