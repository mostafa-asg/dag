package main

import (
	"github.com/mostafa-asg/dag"
)

func main() {

	d := dag.New()
	d.Pipeline(f1, f2).OnComplete(f3).
		Then().
		Spawns(f1, f2).OnComplete(f4)
	d.Run()

}

func f1() {
	println("f1")
}
func f2() {
	println("f2")
}
func f3() {
	println("complete")
}
func f4() {
	println("finish")
}
