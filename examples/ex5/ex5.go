package main

import "github.com/mostafa-asg/dag"

func main() {

	d := dag.New()
	d.Spawns(f1, f2, f3).
		Join().
		Pipeline(f4, f5).
		Then().
		Spawns(f6, f7, f8)
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
func f4() {
	println("f4")
}
func f5() {
	println("f5")
}
func f6() {
	println("f6")
}
func f7() {
	println("f7")
}
func f8() {
	println("f8")
}
