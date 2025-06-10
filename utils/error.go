package utils

import (
	"fmt"
	"log"
)

func Catch(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Panic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
