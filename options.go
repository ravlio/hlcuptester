package hlcuptester

import "strings"
import "regexp"

type Filter func(u string) bool

func InsidePath(s string) Filter {
	return func(u string) bool {
		return  strings.Contains(u, s)
	}
}

func EndOfPath(s string) Filter {
	return func(u string) bool {
		p:=strings.Split(u, "?")
		return strings.HasSuffix(p[0],s)
	}
}

func PathRegexp(rx *regexp.Regexp) Filter {
	return func(u string) bool {
		p:=strings.Split(u, "?")
		return rx.MatchString(p[0])

		return true
	}
}