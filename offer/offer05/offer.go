package offer05

func replaceSpace1(s string) string {
	result := make([]rune, len(s)*3)

	i := 0
	for _, v := range s {
		if v != ' ' {
			result[i] = v
			i++
		} else {
			result[i] = '%'
			result[i+1] = '2'
			result[i+2] = '0'
			i += 3
		}
	}

	return string(result)[:i]
}

func replaceSpace2(s string) string {
	var b []rune
	for _, v := range s {
		if v == 32 {
			b = append(b, 37, 50, 48)
		} else {
			b = append(b, v)
		}
	}
	return string(b)
}
