package main

import (
	"fmt"

	"github.com/kbubj/Go/greetings"
)

func main() {
	massage := greetings.Hello("kbubj")
	fmt.Println(massage)
}
