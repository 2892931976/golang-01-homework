package main

import(
	"fmt"
)

type Student struct{
	Id 		int
	Name 	string

}

func main(){

	stu01 := new(Student)
	stu01.Id = 1
	stu01.Name = "bingan"
	stu02 := new(Student)
	stu02.Id = 2
	stu02.Name = "anc"
	fmt.Println(stu02)


}
