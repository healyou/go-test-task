package utils

type Stack struct {
	s []string
}

func NewFromArray(items []string) Stack {
	return Stack{items}
}

func (s Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) Push(v string) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() string {
	l := len(s.s)
	if l == 0 {
		// опустим проверка на 0, т.к. пока это не требуется -> isEmpty
		return ""
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}
