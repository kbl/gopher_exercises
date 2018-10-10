package ex510

func TopoSortNonDeterministic(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				tempPrereqs := make(map[string]bool)
				for _, prereq := range m[item] {
					tempPrereqs[prereq] = false
				}
				visitAll(tempPrereqs)
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = false
	}

	visitAll(keys)

	return order
}
