package dag

func runSync(job *Job) {

	for _, task := range job.tasks {
		task()
	}
	if job.onComplete != nil {
		job.onComplete()
	}

}
