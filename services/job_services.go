package services

import (
	"fmt"

	"job-queue-system/models"
	"job-queue-system/queue"
	"job-queue-system/repository"
	"go.uber.org/zap"
)

type JobService interface {
	SubmitJob(payload string) (*models.Job, error)
	ProcessJob(job *models.Job)
	GetJob(id int64) (*models.Job, error)
	ListJobs(offset, limit int) ([]models.Job, error)
}

type jobService struct {
	repo   repository.JobRepository
	logger *zap.SugaredLogger
}

func NewJobService(r repository.JobRepository, l *zap.SugaredLogger) JobService {
	return &jobService{r, l}
}

func (s *jobService) SubmitJob(payload string) (*models.Job, error) {
	job := &models.Job{Payload: payload, Status: "queued"}
	err := s.repo.Create(job)
	if err != nil {
		return nil, err
	}
	s.logger.Infow("Job submitted", "id", job.ID)
	queue.AddJobToQueue(job) // No circular import anymore
	return job, nil
}

func (s *jobService) ProcessJob(job *models.Job) {
	s.logger.Infow("Processing job", "id", job.ID)
	job.Status = "processing"
	s.repo.Update(job)
	job.Result = fmt.Sprintf("Processed: %s", job.Payload)
	job.Status = "completed"
	s.repo.Update(job)
	s.logger.Infow("Job completed", "id", job.ID)
}

func (s *jobService) GetJob(id int64) (*models.Job, error) {
	return s.repo.FindByID(id)
}

func (s *jobService) ListJobs(offset, limit int) ([]models.Job, error) {
	return s.repo.List(offset, limit)
}
