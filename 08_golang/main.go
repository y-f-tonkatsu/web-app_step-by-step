package main

import "log"

func main() {
	log.Println("Hello World from go")
	funcSample()
	funcString("passed argment")
}

func funcSample() {
	log.Println("Hello from funcSample")
}

func funcString(word string) {
	log.Println(word, "from funcString")
}
