package main

func filter(results []result) []result {

	m := make(map[string]result)

	for _, r := range results {
		key := r.needle.name + "_IN_" + r.haystack.name
		if v, ok := m[key]; ok {
			if v.avgDiff > r.avgDiff {
				m[key] = r
			}
		} else {
			m[key] = r
		}
	}

	var filtered []result
	for _, v := range m {
		filtered = append(filtered, v)
	}

	return filtered
}
