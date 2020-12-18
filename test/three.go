package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	max := 0
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		cIndex := strings.LastIndex(sb.String(),string(s[i]))
		if cIndex > -1 {
			if max < len(sb.String()){
				max = len(sb.String())
			}
			var str string
			if cIndex < len(sb.String())-1 || cIndex == 0{
				str = sb.String()[cIndex+1:]
			}
			sb.Reset()
			sb.WriteString(str)
		}
		sb.WriteString(string(s[i]))
		if i == len(s) - 1 && max < len(sb.String()){
			max = len(sb.String())
		}
	}
	return max
}

func lengthOfLongestSubstring_(s string) int {
	max := 0
	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		cIndex := strings.LastIndex(s[start:end],string(s[i]))
		if cIndex > -1 {
			if max < (end - start){
				max = end - start
			}
			if cIndex == 0{
				start ++
			} else if cIndex < end - start {
				start += cIndex + 1
			} else {
				start = end
			}
		}
		end ++
		if i == len(s) - 1 && max < end - start{
			max = end - start
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring_("abcabcbb"))
//tt()
}

func tt() {
	a := "abcd"
	fmt.Println(a[1:2])
}
