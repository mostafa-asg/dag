package main

import "github.com/mostafa-asg/dag"

func main() {

	d := dag.New()
	d.Pipeline(f1, f2, f3)
	d.Run()
}

func f1() {
	println("f1")
}
func f2() {
	println("f2")
}
func f3() {
	println("f3")
}
