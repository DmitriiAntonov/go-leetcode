package problem71

import "strings"

func simplifyPath(path string) string {
	parts, stack := strings.Split(path, "/"), Stack{}

	for _, part := range parts {
		switch part {
		case "..":
			if stack.Len() != 0 {
				stack.Pop()
			}
		default:
			if part != "." && part != "" {
				stack.Push(part)
			}
		}
	}

	if stack.Len() == 0 {
		return "/"
	}

	sb := strings.Builder{}

	for _, component := range stack.Slice() {
		sb.WriteString("/")
		sb.WriteString(component)
	}

	return sb.String()
}

// Stack is implemented stack data structure
type Stack struct {
	data []string
}

// Push method add value to the stack
func (s *Stack) Push(value string) {
	s.data = append(s.data, value)
}

// Pop method returns a last added value and removes it from the stack
func (s *Stack) Pop() string {
	if s.Len() == 0 {
		panic("The stack is empty")
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value
}

// Len method returns a length of the internal slice
func (s *Stack) Len() int {
	return len(s.data)
}

// Slice method returns the internal slice
func (s *Stack) Slice() []string {
	return s.data
}
