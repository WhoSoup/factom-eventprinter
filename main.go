package main

import (
	"fmt"

	eater "github.com/WhoSoup/factom-eater"
)

func main() {
	eat, err := eater.Launch(":8040")
	if err != nil {
		panic(err)
	}

	for ev := range eat.Reader() {
		fmt.Println(ev)
	}
}
