package llt

import "strings"

//go:generate mockgen -destination=./mock/mock_human.go -package=mock -source=interface.go
type Human interface {
	Speak() string
	Eat() string
	Get(string, int) string
}

type Animal interface {
	Run() string
}

func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
