package leetcode

func reverseBits(num int) int {
	res, tmp1, tmp2 := 0, 0, 0
	tmp := uint32(num)
	for {
		if tmp&1 == 1 {
			tmp1++
		} else {
			res = max(1+tmp1+tmp2, res)
			if tmp == 0 {
				break
			}
			tmp2 = tmp1
			tmp1 = 0
		}
		tmp >>= 1
	}
	if res > 32 {
		return 32
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
