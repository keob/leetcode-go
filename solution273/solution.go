package solution273

var (
	BILLION  = 1000000000
	MILLION  = 1000000
	THOUSAND = 1000
	HUNDRED  = 100
	TEN      = 10
)

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	billion := num / BILLION
	million := (num - billion*BILLION) / MILLION
	thousand := (num - billion*BILLION - million*MILLION) / THOUSAND
	left := num - billion*BILLION - million*MILLION - thousand*THOUSAND
	ret := ""

	if billion != 0 {
		ret += three(billion) + " Billion"
	}
	if million != 0 {
		if ret != "" {
			ret += " "
		}
		ret += three(million) + " Million"
	}
	if thousand != 0 {
		if ret != "" {
			ret += " "
		}
		ret += three(thousand) + " Thousand"
	}
	if left != 0 {
		if ret != "" {
			ret += " "
		}
		ret += three(left)
	}

	return ret
}

func three(num int) string {
	hundred := num / HUNDRED
	left := num - hundred*HUNDRED

	if hundred == 0 {
		return two(num)
	}

	ret := one(hundred) + " Hundred"

	if left != 0 {
		ret += " " + two(left)
	}

	return ret
}

func two(num int) string {
	if num == 0 {
		return ""
	} else if num < 10 {
		return one(num)
	} else if num < 20 {
		return teen(num)
	}
	ten := num / TEN
	left := num - ten*TEN
	ret := ty(ten)
	if left != 0 {
		ret += " " + one(left)
	}
	return ret
}

func one(num int) string {
	switch num {
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	}

	return ""
}

func teen(num int) string {
	switch num {
	case 10:
		return "Ten"
	case 11:
		return "Eleven"
	case 12:
		return "Twelve"
	case 13:
		return "Thirteen"
	case 14:
		return "Fourteen"
	case 15:
		return "Fifteen"
	case 16:
		return "Sixteen"
	case 17:
		return "Seventeen"
	case 18:
		return "Eighteen"
	case 19:
		return "Nineteen"
	}

	return ""
}

func ty(num int) string {
	switch num {
	case 2:
		return "Twenty"
	case 3:
		return "Thirty"
	case 4:
		return "Forty"
	case 5:
		return "Fifty"
	case 6:
		return "Sixty"
	case 7:
		return "Seventy"
	case 8:
		return "Eighty"
	case 9:
		return "Ninety"
	}

	return ""
}
