// Dependency Inversion: Depender de abstracciones y no de concreciones.
//Las entidades de alto nivel no deben depender de entidades de bajo nivel. Ambas deben depender de abstracciones.
//Las abstracciones no deben depender de los detalles. Los detalles deben depender de las abstracciones.
package main

import "fmt"

// High-level module
type Computer struct {
	monitor Monitor
}

func (c *Computer) Display() {
	c.monitor.Display()
}

// Abstraction
type Monitor interface {
	Display()
}

// Low-level module
type DellMonitor struct{}

func (d *DellMonitor) Display() {
	fmt.Println("Displaying content on Dell monitor")
}

func main() {
	dellMonitor := &DellMonitor{}
	computer := &Computer{monitor: dellMonitor}
	computer.Display()
}