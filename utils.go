package fantasia

import "github.com/nsf/termbox-go"

func ContainsRune(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsTermboxKey(s []termbox.Key, e termbox.Key) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

