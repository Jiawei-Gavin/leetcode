package leetcode

type WordsFrequency struct {
	words map[string]int
}

func Constructor(book []string) WordsFrequency {
	hashMap := make(map[string]int)
	for _, s := range book {
		if _, ok := hashMap[s]; ok {
			hashMap[s] = hashMap[s] + 1
		} else {
			hashMap[s] = 1
		}
	}
	return WordsFrequency{words: hashMap}
}

func (this *WordsFrequency) Get(word string) int {
	return this.words[word]
}
