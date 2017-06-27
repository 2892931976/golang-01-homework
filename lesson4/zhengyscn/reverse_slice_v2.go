package main

func Reverse(s []int) []int {
	length := len(s) - 1
	for i, _ := range s[:length/2] {
		t := s[i]
		s[i] = s[length-i]
		s[length-i] = t
	}
	return s
}

func main() {
	s := []int{2, 3, 5, 7, 11}
	Reverse(s)
	Reverse(s[1:3])

}
