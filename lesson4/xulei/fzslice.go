package main

import (
	"fmt"
	 "os"
	 "strconv"
)

func main() {
	s := []int{2, 3, 5, 7, 11}
	fmt.Println(s)
	if len(os.Args) != 2 {
		fmt.Println("please input lenth:")
		return
        }		      


         lenth, err := strconv.Atoi(os.Args[1])
	 if err != nil {
                 fmt.Println(err)
		 return
         }
         if   len(s) < lenth {
                fmt.Println("fan wei chaochu:")
	 	return
	 } 

	 s1 := s[:lenth]
	// fmt.Println(s1)
	 s2 := s[lenth:]
	 //fmt.Println(s2)
         s2 =append(s2,s1...)
	 fmt.Println(s2)

}
