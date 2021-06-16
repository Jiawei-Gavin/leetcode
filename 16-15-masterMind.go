package leetcode

func masterMind(solution string, guess string) []int {
	res := make([]int, 2)
	hashMap := make(map[byte]int)
	for i := 0; i < len(solution); i++ {
		if solution[i] == guess[i] {
			res[0]++
		}
		hashMap[solution[i]]++
	}
	sum := 0
	for i := 0; i < len(guess); i++ {
		if _, ok := hashMap[guess[i]]; ok && hashMap[guess[i]] > 0 {
			sum++
			hashMap[guess[i]]--
		}
	}
	res[1] = sum - res[0]
	return res
}
