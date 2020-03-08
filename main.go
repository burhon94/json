package main

import "log"

func main() {
	Printer("hi")
}

func Printer(input string) {
	log.Print(input)
}