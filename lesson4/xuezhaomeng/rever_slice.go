package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	slice := []int{1, 3, 5, 7, 9, 11, 13, 15}

	slice2 := slice[num:]
	slice1 := slice[0:num]

	slice2 = append(slice2, slice1...)
        fmt.Printf("反转前的切片为%v\n",slice)
	fmt.Printf("反转后的切片为%v\n",slice2)

}
