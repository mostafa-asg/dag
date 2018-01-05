package pipeline

// Of wraps tasks as a single function
func Of(tasks ...func()) func() {

	return func() {

		for _, task := range tasks {
			task()
		}

	}

}
