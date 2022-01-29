package main

import (
	"fmt"
	demo "github.com/tommyxiang1/gin-go-common/src/demo"
	demo2 "github.com/tommyxiang1/gin-go-common/src/demo2"
)

func main()  {
	fmt.Println(demo.SayHi("Li lei"))
	fmt.Println(demo2.SayHi("Li lei"))

}