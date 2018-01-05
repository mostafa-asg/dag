package dag

import "sync"

func runAsync(job *Job) {

	wg := &sync.WaitGroup{}
	wg.Add(len(job.tasks))

	for _, task := range job.tasks {
		go func(task func()) {
			task()
			wg.Done()
		}(task)
	}

	wg.Wait()
}
