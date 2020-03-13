package main

import (
	"os"
	"strconv"
)

func main() {
	name := os.Args[1]
	num, _ := strconv.Atoi(os.Args[2])

	err := tail(name, num)
	if err != nil{
		panic(err)
	}
}


