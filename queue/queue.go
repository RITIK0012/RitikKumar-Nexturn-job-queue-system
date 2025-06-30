package queue

import "job-queue-system/models"

var jobQueue = make(chan *models.Job, 100)

type JobProcessorFunc func(*models.Job)

func StartWorkerPool(processFunc JobProcessorFunc, count int) {
	for i := 0; i < count; i++ {
		go func() {
			for job := range jobQueue {
				processFunc(job)
			}
		}()
	}
}

func AddJobToQueue(job *models.Job) {
	jobQueue <- job
}
