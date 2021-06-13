package leetcode

func exchangeBits(num int) int {
	even := num & 0xaaaaaaaa
	odd := num & 0x55555555
	return even>>1 | odd<<1
}
