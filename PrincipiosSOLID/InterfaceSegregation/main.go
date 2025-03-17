// Muchas interfaces específicas son mejores que una interfaz general.
package main

import "fmt"

// Printer interface
type Printer interface {
	Print() string
}

// Scanner interface
type Scanner interface {
	Scan() string
}

// MultiFunctionDevice interface
type MultiFunctionDevice interface {
	Printer
	Scanner
}

// Concrete implementation of Printer
type MyPrinter struct{}

func (p MyPrinter) Print() string {
	return "Printing document"
}

// Concrete implementation of Scanner
type MyScanner struct{}

func (s MyScanner) Scan() string {
	return "Scanning document"
}

// Concrete implementation of MultiFunctionDevice
type MyMultiFunctionDevice struct {
	MyPrinter
	MyScanner
}

func main() {
	printer := MyPrinter{}
	scanner := MyScanner{}
	mfd := MyMultiFunctionDevice{}

	fmt.Println(printer.Print())
	fmt.Println(scanner.Scan())
	fmt.Println(mfd.Print())
	fmt.Println(mfd.Scan())
}