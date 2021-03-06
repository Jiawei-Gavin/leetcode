package leetcode

type AnimalShelf struct {
	cat [][]int
	dog [][]int
}

func Constructor() AnimalShelf {
	return AnimalShelf{}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		this.cat = append(this.cat, animal)
	} else if animal[1] == 1 {
		this.dog = append(this.dog, animal)
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if len(this.dog) == 0 && len(this.cat) == 0 {
		return []int{-1, -1}
	}
	if len(this.cat) == 0 {
		res := this.dog[0]
		if len(this.dog) == 1 {
			this.dog = this.dog[:0]
		} else {
			this.dog = this.dog[1:]
		}
		return res
	}
	if len(this.dog) == 0 {
		res := this.cat[0]
		if len(this.cat) == 1 {
			this.cat = this.cat[:0]
		} else {
			this.cat = this.cat[1:]
		}
		return res
	}
	if this.cat[0][0] < this.dog[0][0] {
		res := this.cat[0]
		if len(this.cat) == 1 {
			this.cat = this.cat[:0]
		} else {
			this.cat = this.cat[1:]
		}
		return res
	} else {
		res := this.dog[0]
		if len(this.dog) == 1 {
			this.dog = this.dog[:0]
		} else {
			this.dog = this.dog[1:]
		}
		return res
	}
}

func (this *AnimalShelf) DequeueDog() []int {
	if len(this.dog) == 0 {
		return []int{-1, -1}
	}
	res := this.dog[0]
	if len(this.dog) == 1 {
		this.dog = this.dog[:0]
	} else {
		this.dog = this.dog[1:]
	}
	return res
}

func (this *AnimalShelf) DequeueCat() []int {
	if len(this.cat) == 0 {
		return []int{-1, -1}
	}
	res := this.cat[0]
	if len(this.cat) == 1 {
		this.cat = this.cat[:0]
	} else {
		this.cat = this.cat[1:]
	}
	return res
}
