package solution373

import "sort"

func kSmallestPairs(nums1, nums2 []int, k int) (res [][]int) {
	m, n := len(nums1), len(nums2)

	left, right := nums1[0]+nums2[0], nums1[m-1]+nums2[n-1]+1
	pairSum := left + sort.Search(right-left, func(sum int) bool {
		sum += left
		cnt := 0
		i, j := 0, n-1
		for i < m && j >= 0 {
			if nums1[i]+nums2[j] > sum {
				j--
			} else {
				cnt += j + 1
				i++
			}
		}
		return cnt >= k
	})

	i := n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] >= pairSum {
			i--
		}
		for _, num2 := range nums2[:i+1] {
			res = append(res, []int{num1, num2})
			if len(res) == k {
				return
			}
		}
	}

	i = n - 1
	for _, num1 := range nums1 {
		for i >= 0 && num1+nums2[i] > pairSum {
			i--
		}
		for j := i; j >= 0 && num1+nums2[j] == pairSum; j-- {
			res = append(res, []int{num1, nums2[j]})
			if len(res) == k {
				return
			}
		}
	}

	return
}
