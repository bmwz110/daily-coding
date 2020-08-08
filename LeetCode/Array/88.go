package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Println(countCharacters(words, chars))
}

// 遍历单词数组 -> 遍历单词字母
// if (单词中每个字母的数量 均 <= chars中对应字母数量) { 则该单词可以被拼写 }
func countCharacters(words []string, chars string) int {
	// chars中的只能用一次
	var wholeLength int = 0
loop:
	for _, v1 := range words {
		for _, v2 := range v1 {
			// 取字母
			if strings.Count(v1, string(v2)) > strings.Count(chars, string(v2)) {
				continue loop
			}
		}
		wholeLength += len(v1)
	}
	return wholeLength
} 
 
