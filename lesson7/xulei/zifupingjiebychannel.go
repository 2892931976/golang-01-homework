package main


import "fmt"

func pj(word []string, c chan string) {

	var words string


	for _, v := range word {

		words += v
	}
    c <- words
}

func main() {

	 s := []string{"hello", "golang", "c++", "world"}

	c := make(chan string)
	d := make(chan string)

	go pj(s[len(s)/2:], c)
	go pj(s[:len(s)/2], d)


	x, y := <-c, <-d

	fmt.Println(y, x, x + y)




}
