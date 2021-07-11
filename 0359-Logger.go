package leetcode

type Logger struct {
	timeMap map[string]int
}

func Constructor() Logger {
	return Logger{timeMap: map[string]int{}}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if this.timeMap[message] == 0 {
		this.timeMap[message] = timestamp + 10
		return true
	} else if timestamp >= this.timeMap[message] {
		this.timeMap[message] = timestamp + 10
		return true
	}
	return false
}
