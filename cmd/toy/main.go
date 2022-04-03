package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TOMOFUMI-KONDO/toy/interpreter"
	"github.com/TOMOFUMI-KONDO/toy/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "toy file path must be passed.")
		os.Exit(1)
	}

	path := os.Args[1]
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read file %q: %v", path, err)
	}

	toy := &parser.Toy{Buffer: string(input)}
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
