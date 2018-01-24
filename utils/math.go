package utils

import (
	"strconv"
	"strings"
)

var (
	Tokens string
	Length int
)

func init() {
	//0-9
	for i := 0; i < 10; i++ {
		Tokens += strconv.Itoa(i)
	}
	//a-z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('a') + byte(i))
	}
	//A-Z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('A') + byte(i))
	}
	Length = len(Tokens)
}
func IdToString(i int) string {
	var res string
	for i > 0 {
		m := i % Length
		res = string(Tokens[m]) + res
		i /= Length
	}
	return res
}

//res=
func StringToId(s string) int {
	var res, length int = 0, len(s)
	for k, v := range s {
		if length-k > 1 {
			pro := 1
			for i := 1; i < length-k; i++ {
				pro = pro * Length
			}
			pro = pro * strings.Index(Tokens, string(v))
			res += pro
		}
	}
	return res + strings.Index(Tokens, string(s[length-1]))
}
