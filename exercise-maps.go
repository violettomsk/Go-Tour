package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int);
	ss := strings.Fields(s);
	for _, v := range ss {
		if _, present := wc[v]; present == true {
			wc[v]++;
		} else {
			wc[v] = 1;
		}
	}
	return wc;
}

func main() {
	wc.Test(WordCount)
}
