package leetcode

func isUnique(astr string) bool {
	mark := 0
	for _, s := range astr {
		move := s - 'a'
		if mark&(1<<move) != 0 {
			return false
		} else {
			mark |= 1 << move
		}
	}
	return true
}
