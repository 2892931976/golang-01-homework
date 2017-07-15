package main

import (
	"fmt"
)

type myStruct struct {
	name   bool
	userid int64
}

var structZero myStruct
var intZero int
var int32Zero int32
var int64Zero int64
var uintZero uint
var uint8Zero uint8
var uint32Zero uint32
var uint64Zero uint64
var byteZero byte
var boolZero bool
var float32Zero float32
var float64Zero float64
var stringZero string
var funcZero func(int) int
var byteArrayZero [5]byte
var boolArrayZero [5]bool
var byteSliceZero []byte
var boolSliceZero []bool
var mapZero map[string]bool
var interfaceZero interface{}
var chanZero chan int
var pointerZero *int

func main() {
	fmt.Println("structZero", structZero)
	fmt.Println("intZero", intZero)
	fmt.Println("int32Zero", int32Zero)
	fmt.Println("int64Zero", int64Zero)
	fmt.Println("uintZero", uintZero)
	fmt.Println("unit8Zero", uint8Zero)
	fmt.Println("uint32Zero", uint32Zero)
	fmt.Println("unit64Zero", uint64Zero)
	fmt.Println("byteZero", byteZero)
	fmt.Println("boolZero", boolZero)
	fmt.Println("float64Zero", float64Zero)
	fmt.Println("stringZero", stringZero)
	fmt.Println("funcZero", funcZero)
	fmt.Println("funcZero == nil? ", funcZero == nil)
	fmt.Println("byteArrayZero:", byteArrayZero)
	fmt.Println("boolArrayZero: ", boolArrayZero)
	fmt.Println("byteSliceZero: ", byteSliceZero)
	fmt.Println("boolSliceZero: ", boolSliceZero)
	fmt.Println("mapZero: ", mapZero)
	fmt.Println("mapZero's len ? ", len(mapZero))
	fmt.Println("mapZero == nil: ", mapZero == nil)
	fmt.Println("interfaceZero: ", interfaceZero)
	fmt.Println("interfaceZero == nil?", interfaceZero == nil)
	fmt.Println("chanZero", chanZero)
	fmt.Println("chanZero == nil : ", chanZero)
	fmt.Println("pointerZero : ", pointerZero)
	fmt.Println("pointerZero == nil", pointerZero == nil)

	fmt.Println("byteSliceZero's len ", len(byteSliceZero))
	fmt.Println("byteSliceZero 's cap", cap(byteSliceZero))
	fmt.Println("byteSliceZero == nil ? ", byteSliceZero == nil)

}
