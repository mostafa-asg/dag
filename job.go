package dag

// Job - Each job consists of one or more tasks
type Job struct {
	tasks      []func()
	sequential bool
}
