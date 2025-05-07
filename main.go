package main

import (
	"bada/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("welocome %s to bada programmig lang \n", user.Username)
	fmt.Printf("Write any command \n")
	repl.Start(os.Stdin, os.Stdout)
}
