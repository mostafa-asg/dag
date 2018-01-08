package dag

// Dag represents directed acyclic graph
type Dag struct {
	jobs []*Job
}

// New creates new DAG
func New() *Dag {
	return &Dag{
		jobs: make([]*Job, 0),
	}
}

func (dag *Dag) lastJob() *Job {
	jobsCount := len(dag.jobs)
	if jobsCount == 0 {
		return nil
	}

	return dag.jobs[jobsCount-1]
}

// Run starts the tasks
// It will block until all functions are done
func (dag *Dag) Run() {

	for _, job := range dag.jobs {
		run(job)
	}

}

// RunAsync executes Run on another goroutine
func (dag *Dag) RunAsync(onComplete func()) {
	go func() {

		dag.Run()

		if onComplete != nil {
			onComplete()
		}

	}()
}

// Pipeline executes tasks sequentially
func (dag *Dag) Pipeline(tasks ...func()) *pipelineResult {

	job := &Job{
		tasks:      make([]func(), len(tasks)),
		sequential: true,
	}

	for i, task := range tasks {
		job.tasks[i] = task
	}

	dag.jobs = append(dag.jobs, job)

	return &pipelineResult{
		dag,
	}
}

// Spawns executes tasks concurrently
func (dag *Dag) Spawns(tasks ...func()) *spawnsResult {

	job := &Job{
		tasks:      make([]func(), len(tasks)),
		sequential: false,
	}

	for i, task := range tasks {
		job.tasks[i] = task
	}

	dag.jobs = append(dag.jobs, job)

	return &spawnsResult{
		dag,
	}
}
