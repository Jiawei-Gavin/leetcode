package leetcode

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

// solution1
func getImportance(employees []*Employee, id int) int {
	mp, total := make(map[int]int, 0), 0
	for idx, employee := range employees {
		mp[employee.Id] = idx
	}

	var dfs func(int)
	dfs = func(id int) {
		employee := employees[id]
		total += employee.Importance
		for _, subId := range employee.Subordinates {
			dfs(mp[subId])
		}
	}
	dfs(mp[id])
	return total
}

// solution2
func getImportance(employees []*Employee, id int) int {
	mp, queue, total := map[int]*Employee{}, []int{id}, 0
	for _, employee := range employees {
		mp[employee.Id] = employee
	}

	for len(queue) > 0 {
		employee := mp[queue[0]]
		queue = queue[1:]
		if employee == nil {
			continue
		}
		total += employee.Importance
		for _, subId := range employee.Subordinates {
			queue = append(queue, subId)
		}
	}
	return total
}
