package leetcode

var wordMap = map[int]string{
	0:          "Zero",
	1:          "One",
	2:          "Two",
	3:          "Three",
	4:          "Four",
	5:          "Five",
	6:          "Six",
	7:          "Seven",
	8:          "Eight",
	9:          "Nine",
	10:         "Ten",
	11:         "Eleven",
	12:         "Twelve",
	13:         "Thirteen",
	14:         "Fourteen",
	15:         "Fifteen",
	16:         "Sixteen",
	17:         "Seventeen",
	18:         "Eighteen",
	19:         "Nineteen",
	20:         "Twenty",
	30:         "Thirty",
	40:         "Forty",
	50:         "Fifty",
	60:         "Sixty",
	70:         "Seventy",
	80:         "Eighty",
	90:         "Ninety",
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}

func numberToWords(num int) string {
	if num == 0 {
		return wordMap[0]
	}
	result := ""
	quotient := num / 1000000000
	num = num % 1000000000
	if quotient > 0 {
		result += ParseNumberLessThanOneThousand(quotient) + " " + "Billion"
	}
	quotient = num / 1000000
	num = num % 1000000
	if quotient > 0 {
		if result != "" {
			result += " "
		}
		result += ParseNumberLessThanOneThousand(quotient) + " " + "Million"
	}
	quotient = num / 1000
	num = num % 1000
	if quotient > 0 {
		if result != "" {
			result += " "
		}
		result += ParseNumberLessThanOneThousand(quotient) + " " + "Thousand"
	}
	if num > 0 {
		if result != "" {
			result += " "
		}
		result += ParseNumberLessThanOneThousand(num)
	}
	return result
}

func ParseNumberLessThanOneThousand(num int) string {
	result := ""
	quotient := num / 100
	num = num % 100
	if quotient != 0 {
		result += wordMap[quotient] + " " + "Hundred"
	}
	if num == 0 {
		return result
	}
	if num >= 20 {
		remainder := num % 10
		if remainder > 0 {
			if result != "" {
				result += " "
			}
			result += wordMap[num-remainder] + " " + wordMap[remainder]
		} else {
			if result != "" {
				result += " "
			}
			result += wordMap[num-remainder]
		}
	} else if num > 0 && num < 20 {
		if result != "" {
			result += " "
		}
		result += wordMap[num]
	}
	return result
}
