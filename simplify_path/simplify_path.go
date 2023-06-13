package simplify_path

import "strings"

func simplifyPath(path string) string {
	parts, stack := strings.Split(path, "/"), Stack{}

	for _, part := range parts {
		switch part {
		case ".", "":
		// Do nothing
		case "..":
			if stack.Len() != 0 {
				stack.Pop()
			}
		default:
			stack.Push(part)

		}
	}

	if stack.Len() == 0 {
		return "/"
	}

	sb := strings.Builder{}

	for _, component := range stack.data {
		sb.WriteString("/")
		sb.WriteString(component)
	}

	return sb.String()
}

type Stack struct {
	data []string
}

func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() string {
	if s.Len() == 0 {
		panic("The stack is empty")
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Slice() []string {
	return s.data
}
