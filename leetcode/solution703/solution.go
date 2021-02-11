package solution703

type KthLargest struct {
	size  int
	data  []int
	count int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{}
	kth.size = k
	kth.data = []int{0}

	for _, num := range nums {
		kth.Add(num)
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	if this.count < this.size-1 {
		this.data = append(this.data, val)
		this.count += 1
	} else if this.count == this.size-1 {
		this.data = append(this.data, val)
		this.count += 1
		n := len(this.data) - 1
		for i := n / 2; i >= 1; i-- {
			heapify(this.data, i)
		}
	} else {
		if val > this.data[1] {
			this.data[1] = val
			heapify(this.data, 1)
		}
	}
	return this.data[1]
}

func heapify(a []int, i int) {
	n := len(a) - 1
	for {
		minPos := i
		if i*2 <= n && a[i*2] < a[minPos] {
			minPos = i * 2
		}
		if i*2+1 <= n && a[i*2+1] < a[minPos] {
			minPos = i*2 + 1
		}
		if minPos == i {
			break
		}
		a[minPos], a[i] = a[i], a[minPos]
		i = minPos
	}
}
