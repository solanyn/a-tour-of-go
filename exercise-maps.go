package main

import (
	"golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    tokens := strings.Fields(s)
    m := map[string]int{}
    for _, t := range tokens {
        _, ok := m[t]
        if ok {
            m[t]++
        } else {
            m[t] = 1
        }
    }
	return m
}

func main() {
	wc.Test(WordCount)
}

