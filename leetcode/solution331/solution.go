package solution331

import "strings"

func isValidSerialization(preorder string) bool {
	nodeArray := strings.Split(preorder, ",")
	slot := 1

	for i := 0; i < len(nodeArray); i++ {
		slot -= 1
		if slot < 0 {
			return false
		}

		if nodeArray[i] != "#" {
			slot += 2
		}
	}

	return slot == 0
}
