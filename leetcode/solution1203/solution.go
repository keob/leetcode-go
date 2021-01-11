package solution1203

func sortItems(n int, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n)
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n)
	itemDegree := make([]int, n)

	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID {
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else {
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}

	groupOrders := sort(groupGraph, groupDegree, items)

	if len(groupOrders) < len(items) {
		return nil
	}

	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := sort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}

	return
}

func sort(graph [][]int, deg []int, items []int) (res []int) {
	tmp := []int{}

	for _, i := range items {
		if deg[i] == 0 {
			tmp = append(tmp, i)
		}
	}

	for len(tmp) > 0 {
		from := tmp[0]
		tmp = tmp[1:]

		res = append(res, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				tmp = append(tmp, to)
			}
		}
	}
	return
}
