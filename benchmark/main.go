package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/florianwoelki/reflow/compiler"
	"github.com/florianwoelki/reflow/evaluator"
	"github.com/florianwoelki/reflow/lexer"
	"github.com/florianwoelki/reflow/object"
	"github.com/florianwoelki/reflow/parser"
	"github.com/florianwoelki/reflow/vm"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

var input = `
let fibo = fn(x) {
	if (x == 0) {
		0
	} else {
		if (x == 1) {
			1
		} else {
			fibo(x - 1) + fibo(x - 2)
		}
	}
}

fibo(35);
`

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	if *engine == "vm" {
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.Bytecode())

		start := time.Now()

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPoppedStackElem()
	} else {
		env := object.NewEnvironment()
		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)
	}

	fmt.Printf("engine=%s, result=%s, duration=%s\n", *engine, result.Inspect(), duration)
}
