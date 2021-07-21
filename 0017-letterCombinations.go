package leetcode

var phoneMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) == 0 {
		return res
	}
	backtrack("", digits, 0)
	return res
}

func backtrack(combination string, digits string, index int) {
	if index == len(digits) {
		res = append(res, combination)
	} else {
		digit := digits[index]
		letters := phoneMap[string(digit)]
		for _, letter := range letters {
			combination += letter
			backtrack(combination, digits, index+1)
			combination = combination[:index]
		}
	}
}
