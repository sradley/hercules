package pool

var done bool

type Result struct {
	Username string
	Password string
	Success  bool
}

type Job interface {
	Run() (*Result, error)
}

func Start(jobs []Job, threads int) {
	done = false

	jobs_ch := make(chan Job, len(jobs))
	resu_ch := make(chan *Result, len(jobs))

	for i := 1; i <= threads; i++ {
		go worker(jobs_ch, resu_ch)
	}

	for _, j := range jobs {
		if !done {
			jobs_ch <- j
		}
	}
	close(jobs_ch)

	for _ = range jobs {
		if !done {
			<- resu_ch
		}
	}
}

func worker(jobs <- chan Job, results chan <- *Result) {
	for j := range jobs {
		if done {
			break
		}

		result, err := j.Run()
		if err != nil {
			panic(err)
		}

		if result.Success {
			// kind of a race condition, sorry lmao
			done = true
			close(results)
		}

		if !done {
			results <- result
		}
	}
}
