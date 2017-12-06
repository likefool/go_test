package main

import (
	"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	wc_map := make(map[string]int)
	str_arr := strings.Fields(s)
	for k, v := range str_arr {        
		fmt.Println("k:", k, "v:", v)
		v_word_count, ok := wc_map[v]
		if ok == false {
			wc_map[v] = 1
		} else {
			wc_map[v] = v_word_count + 1
		}
	}
	//return map[string]int{"x": 1}
	return wc_map
}

func main() {
	wc.Test(WordCount)
}

