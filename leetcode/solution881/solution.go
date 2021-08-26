package solution881

import (
	"sort"
)

func numRescueBoats(people []int, limit int) (res int) {
	sort.Ints(people)
	light, heavy := 0, len(people)-1

	for light <= heavy {
		if people[light]+people[heavy] > limit {
			heavy--
		} else {
			light++
			heavy--
		}
		res++
	}

	return
}
