package main

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Len() int
	IsEmpty() bool
}

type StringStack struct {
	data []string
}

func (s *StringStack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *StringStack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *StringStack) Len() int {
	return len(s.data)
}

func (s *StringStack) IsEmpty() bool {
	return len(s.data) == 0
}

type IntStack struct {
	data []int
}

func (s *IntStack) Push(v int) {
	s.data = append(s.data, v)
}

func (s *IntStack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *IntStack) Len() int {
	return len(s.data)
}

func (s *IntStack) IsEmpty() bool {
	return len(s.data) == 0
}
