package dag

// Job - Each job consists of one or more tasks
// Each Job can runs tasks in order(Sequential) or unordered
type Job struct {
	tasks      []func()
	sequential bool
	onComplete func()
}
