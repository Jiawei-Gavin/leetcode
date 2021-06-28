package leetcode

func convertToTitle(columnNumber int) string {
	if columnNumber <= 26 {
		return string('A' + rune(columnNumber-1)%26)
	}
	return convertToTitle((columnNumber-1)/26) + convertToTitle((columnNumber-1)%26+1)
}
