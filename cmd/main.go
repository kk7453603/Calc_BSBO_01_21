package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	var strFirst, strSecond string

	fmt.Scanln(&strFirst)
	fmt.Scanln(&strSecond)

	first, err := strconv.Atoi(strFirst)
	if err != nil {
		log.Println(err)
	}
	second, err := strconv.Atoi(strSecond)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(first, second)
}
