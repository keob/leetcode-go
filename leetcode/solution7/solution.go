package solution7

func reverse(x int) (res int) {
	for x != 0 {
		if tmp := int32(res); (tmp*10)/10 != tmp {
			return 0
		}
		res = res*10 + x%10
		x = x / 10
	}

	return
}
