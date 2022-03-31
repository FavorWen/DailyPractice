package main

import "fmt"

func backspace(s string, i int) string {
	if i == 0 {
		return s[1:]
	} else if i == 1 {
		return s[2:]
	}
	return s[:i-1] + s[i+1:]
}

func backspaceCompare(s string, t string) (bool, string, string) {
	for i := 0; i < len(s); {
		if s[i] == '#' {
			s = backspace(s, i)
			if i != 0 {
				i--
			}
		} else {
			i++
		}
	}
	for i := 0; i < len(t); {
		if t[i] == '#' {
			t = backspace(t, i)
			if i != 0 {
				i--
			}
		} else {
			i++
		}
	}
	return s == t, s, t
}

func main() {
	s := "ab##"
	t := "c#d#"
	_, s, t = backspaceCompare(s, t)
	fmt.Println(s, t)
}
