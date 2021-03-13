package solution705

type MyHashSet struct {
	bitset []uint64
}

func Constructor() MyHashSet {
	return MyHashSet{bitset: []uint64{}}
}

func (set *MyHashSet) Add(key int) {
	bit := key % 64
	length := key / 64
	for i := len(set.bitset); i <= length; i++ {
		set.bitset = append(set.bitset, 0)
	}
	set.bitset[length] = set.bitset[length] | (1 << bit)
}

func (set *MyHashSet) Remove(key int) {
	bit := key % 64
	length := key / 64
	if length >= len(set.bitset) {
		return
	}
	set.bitset[length] = set.bitset[length] & ^(1 << bit)
}

func (set *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key / 64
	if length >= len(set.bitset) {
		return false
	}
	return set.bitset[length]&(1<<bit) != 0
}
