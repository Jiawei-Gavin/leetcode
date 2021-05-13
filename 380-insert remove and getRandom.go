package leetcode

import "math/rand"

type RandomizedSet struct {
	dict map[int]int
	list []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]int), make([]int, 0)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}
	this.list = append(this.list, val)
	this.dict[val] = len(this.list) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dict[val]; !ok {
		return false
	}
	index := this.dict[val]
	this.list[index] = this.list[len(this.list)-1]
	this.dict[this.list[index]] = index
	this.list = this.list[:len(this.list)-1]
	delete(this.dict, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.list) == 0 {
		return -1
	}
	index := rand.Intn(len(this.list))
	return this.list[index]
}
