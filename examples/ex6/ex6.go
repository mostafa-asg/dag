package main

import "github.com/mostafa-asg/dag"
import "github.com/mostafa-asg/dag/pipeline"

func main() {

	d := dag.New()
	d.Spawns(pipeline.Of(f1, f3), pipeline.Of(f2, f4)).
		Join().
		Pipeline(f5)
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
