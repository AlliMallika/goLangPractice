package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hai! What's your name?")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("Hai!", name)
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hai!", name)
}
