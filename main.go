package main

import (
	"fmt"
	"log"
	"panicrecover/panicrecover"
)

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic Occurred: ", err)
		}
	}()
	fmt.Println(panicrecover.Divide(1, 0))
}

func main() {
	divideByZero()
	fmt.Println("Survived dividing number by 0!")
}
