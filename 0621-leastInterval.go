package leetcode

func leastInterval(tasks []byte, n int) int {
	hashMap := make(map[byte]int)
	for _, task := range tasks {
		hashMap[task]++
	}

	maxExec, maxExecCnt := 0, 0
	for _, val := range hashMap {
		if val > maxExec {
			maxExec, maxExecCnt = val, 1
		} else if val == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}
