package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"a", ""}
	fmt.Println(longestWord(words))
}

func longestWord(words []string) string {
	myMap := make(map[string]int)
	for _, w := range words {
		myMap[w] = len(w)
	}

	ret := ""

	ok := false
	for _, w := range words {
		for i := len(w); i > 0; i-- {
			s := w[0:i]
			if _, ok = myMap[s]; !ok {
				break
			}
		}
		if ok {
			if ret == "" || len(w) > len(ret) {
				ret = w
			} else if len(w) == len(ret) {
				s := sort.StringSlice{w, ret}
				sort.Sort(s)
				ret = s[0]
			}
		}
	}
	return ret
}
