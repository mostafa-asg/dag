package main

import (
	"sync"

	"github.com/mostafa-asg/dag"
)

var wg = &sync.WaitGroup{}

func main() {

	wg.Add(1)

	d := dag.New()
	d.Pipeline(f1, f2).Then().Spawns(f3, f4)
	d.RunAsync(onComplete)

	wg.Wait()
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
func onComplete() {
	println("All functions have completed")
	wg.Done()
}
