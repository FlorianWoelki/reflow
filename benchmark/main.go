package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/florianwoelki/reflow/compiler"
	"github.com/florianwoelki/reflow/evaluator"
	"github.com/florianwoelki/reflow/lexer"
	"github.com/florianwoelki/reflow/object"
	"github.com/florianwoelki/reflow/parser"
	"github.com/florianwoelki/reflow/vm"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")
var filename = flag.String("filename", "benchmark/fibo.rf", "specify filename with '.rf' extension")

func readFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	input, err := readFile(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}

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
