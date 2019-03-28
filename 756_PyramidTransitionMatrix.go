package main

func helper(bottom string, k string, i int, dict map[string]string) bool {
	if len(bottom) == 1 {
		return true
	}
	if i == len(bottom)-1 {
		return helper(k, "", 0, dict)
	} else {
		for _, c := range dict[bottom[i:i+2]] {
			newK := string(append([]rune(k), c))
			if helper(bottom, newK, i+1, dict) {
				return true
			}
		}
	}
	return false
}

func pyramidTransition(bottom string, allowed []string) bool {
	dict := make(map[string]string)
	for _, allow := range allowed {
		dict[allow[:2]] += allow[2:]
	}

	return helper(bottom, "", 0, dict)
}
