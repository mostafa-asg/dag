package dag

func run(job *Job) {

	if job.sequential {
		runSync(job)
	} else {
		runAsync(job)
	}

}
