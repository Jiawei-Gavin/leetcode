package leetcode

import (
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	ss := strings.ToUpper(strings.ReplaceAll(s, "-", ""))
	var res []byte
	if len(ss)%k != 0 {
		res = append(append(res, ss[:len(ss)%k]...), '-')
	}
	for i := len(ss) % k; i < len(ss); i += k {
		res = append(append(res, ss[i:i+k]...), '-')
	}
	if len(res) > 0 && res[len(res)-1] == '-' {
		res = res[:len(res)-1]
	}
	return string(res)
}
