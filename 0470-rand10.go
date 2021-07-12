package leetcode

func rand10() int {
	num := (rand7()-1)*7 + rand7()
	for num > 10 {
		num = (rand7()-1)*7 + rand7()
	}
	return num
}
