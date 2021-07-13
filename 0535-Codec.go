package leetcode

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

type Codec struct {
	hashMap map[int]string
	key     int
}

func Constructor() Codec {
	return Codec{
		hashMap: make(map[int]string),
		key:     rand.Intn(math.MaxInt32),
	}
}

func (this *Codec) encode(longUrl string) string {
	for this.hashMap[this.key] != "" {
		this.key = rand.Intn(math.MaxInt32)
	}
	this.hashMap[this.key] = longUrl
	return "http://tinyurl.com/" + strconv.Itoa(this.key)
}

func (this *Codec) decode(shortUrl string) string {
	temp, _ := strconv.Atoi(strings.Replace(shortUrl, "http://tinyurl.com/", "", -1))
	return this.hashMap[temp]
}
