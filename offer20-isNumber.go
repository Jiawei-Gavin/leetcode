package leetcode

var states = []map[byte]int{
	{' ': 0, 's': 1, 'd': 2, '.': 4},
	{'d': 2, '.': 4},
	{'d': 2, '.': 3, 'e': 5, ' ': 8},
	{'d': 3, 'e': 5, ' ': 8},
	{'d': 3},
	{'s': 6, 'd': 7},
	{'d': 7},
	{'d': 7, ' ': 8},
	{' ': 8},
}

func isNumber(s string) bool {
	p := 0
	var t byte
	for _, v := range s {
		if v >= '0' && v <= '9' {
			t = 'd'
		} else if v == '+' || v == '-' {
			t = 's'
		} else if v == 'e' || v == 'E' {
			t = 'e'
		} else if v == '.' || v == ' ' {
			t = byte(v)
		} else {
			t = '?'
		}
		if _, ok := states[p][t]; !ok {
			return false
		}
		p = states[p][t]
	}
	return p == 2 || p == 3 || p == 7 || p == 8
}
