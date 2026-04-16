package concurrency

import "sync"

type Job func() error

func RunUploaderPool(workerCount int, jobs []Job) error {
	if workerCount <= 0 {
		workerCount = 1
	}

	jobCh := make(chan Job)
	errCh := make(chan error, len(jobs))
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobCh {
				if err := job(); err != nil {
					errCh <- err
				}
			}
		}()
	}

	for _, job := range jobs {
		jobCh <- job
	}
	close(jobCh)

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
