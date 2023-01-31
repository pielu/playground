package main

import (
	"flag"
	"fmt"
	"log"
	"main/roman"
	"strconv"
)

func main() {
	var action string
	var input string
	var program string

	flag.StringVar(&program, "p", "roman", "Specify program. Available: roman.")
	flag.StringVar(&action, "a", "toInt", "Specify action. Depends on the program.")
	flag.StringVar(&input, "i", "", "Parameter for the action.")

	flag.Parse()

	if program == "roman" && action == "toInt" {
		fmt.Println(roman.RomanToInt(input))
	}

	if program == "roman" && action == "intToRoman" {
		si, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(roman.IntToRoman(si))
	}
}
