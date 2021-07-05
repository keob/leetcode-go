package solution726

import (
	"sort"
	"strconv"
)

func isDigit(s byte) bool       { return s >= '0' && s <= '9' }
func isLowerLetter(s byte) bool { return s >= 'a' && s <= 'z' }
func isUpperLetter(s byte) bool { return s >= 'A' && s <= 'Z' }

func countOfAtoms(formula string) string {
	strStack := []string{}
	numStack := []int{}

	var pre byte
	for i := 0; i < len(formula); i++ {
		s := formula[i]
		haveNext := i+1 < len(formula)

		if isUpperLetter(s) {
			strStack = append(strStack, string(s))
			if (haveNext && !isDigit(formula[i+1]) && !isLowerLetter(formula[i+1])) || !haveNext {
				numStack = append(numStack, 1)
			}
			pre = s
			continue
		}

		if isLowerLetter(s) {
			strStack[len(strStack)-1] = strStack[len(strStack)-1] + string(s)
			if (haveNext && !isDigit(formula[i+1]) && !isLowerLetter(formula[i+1])) || !haveNext {
				numStack = append(numStack, 1)
			}
			pre = s
			continue
		}

		if isDigit(s) {
			if isDigit(pre) {
				numStr := strconv.Itoa(numStack[len(numStack)-1])
				numStr += string(s)

				newNum, _ := strconv.Atoi(numStr)
				numStack[len(numStack)-1] = newNum
			} else {
				num, _ := strconv.Atoi(string(s))
				numStack = append(numStack, num)
			}
			pre = s
			continue
		}

		if string(s) == ")" {
			ss := []string{}
			ns := []int{}

			base := 1
			if i+1 < len(formula) && isDigit(formula[i+1]) {
				j := 1
				numStr := ""
				for i+1 < len(formula) && isDigit(formula[i+1]) {
					if i+j <= len(formula)-1 && isDigit(formula[i+j]) {
						numStr += string(formula[i+j])
						j++
					} else {
						break
					}
				}
				base, _ = strconv.Atoi(numStr)
				i = i + (j - 1)
			}

			for {
				str := strStack[len(strStack)-1]
				strStack = strStack[:len(strStack)-1]

				if str == "(" {
					break
				}

				ss = append(ss, str)

				num := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				ns = append(ns, num)
			}

			ln := len(ns)
			for i := 0; i < ln; i++ {
				num := ns[len(ns)-1]
				ns = ns[:len(ns)-1]

				newNum := num * base
				numStack = append(numStack, newNum)

				str := ss[len(ss)-1]
				ss = ss[:len(ss)-1]

				strStack = append(strStack, str)
			}
			pre = s
			continue
		}

		strStack = append(strStack, string(s))
		pre = s
	}

	m := map[string]int{}
	atoms := []string{}
	for i := 0; i < len(strStack); i++ {
		m[strStack[i]] = m[strStack[i]] + numStack[i]
		atoms = add(atoms, strStack[i])
	}

	sort.Strings(atoms)

	res := ""
	for _, atom := range atoms {
		if m[atom] == 1 {
			res += atom
			continue
		}
		res += atom + strconv.Itoa(m[atom])
	}

	return res
}

func add(s []string, e string) []string {
	for _, v := range s {
		if v == e {
			return s
		}
	}

	s = append(s, e)

	return s
}
