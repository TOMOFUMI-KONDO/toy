package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TOMOFUMI-KONDO/toy/interpreter"
	"github.com/TOMOFUMI-KONDO/toy/parser"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed to read stdin: %v", err)
	}

	toy := &parser.Toy{Buffer: string(b)}
	if err := toy.Init(); err != nil {
		panic(err)
	}
	if err := toy.Parse(); err != nil {
		panic(err)
	}
	if err := toy.ConvertAst(); err != nil {
		panic(err)
	}

	itpr := interpreter.NewInterpreter()

	result, err := itpr.CallMain(toy.Program)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
