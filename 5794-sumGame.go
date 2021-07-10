package leetcode

func sumGame(num string) bool {
	sum, cnt := 0, 0
	for i := 0; i < len(num); i++ {
		temp := num[i]
		if i < len(num)/2 {
			if temp == '?' {
				cnt++
			} else {
				sum += int(temp - '0')
			}
		} else {
			if temp == '?' {
				cnt--
			} else {
				sum -= int(temp - '0')
			}
		}
	}
	if cnt%2 == 1 {
		return true
	} else if sum == cnt/2*-9 {
		return false
	} else {
		return true
	}
}
