package dag

func runSync(job *Job) {

	for _, task := range job.tasks {
		task()
	}

}
