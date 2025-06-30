package repository

import (
	"database/sql"
	"time"

	"job-queue-system/models"
)

type JobRepository interface {
	Create(job *models.Job) error
	FindByID(id int64) (*models.Job, error)
	Update(job *models.Job) error
	List(offset, limit int) ([]models.Job, error)
}

type jobRepo struct {
	db *sql.DB
}

func NewJobRepository(db *sql.DB) JobRepository {
	return &jobRepo{db}
}

func (r *jobRepo) Create(job *models.Job) error {
	now := time.Now()
	job.CreatedAt = now
	job.UpdatedAt = now
	query := `INSERT INTO jobs (payload, status, created_at, updated_at) VALUES (?, ?, ?, ?)`
	res, err := r.db.Exec(query, job.Payload, job.Status, job.CreatedAt, job.UpdatedAt)
	if err != nil {
		return err
	}
	job.ID, _ = res.LastInsertId()
	return nil
}

func (r *jobRepo) FindByID(id int64) (*models.Job, error) {
	job := &models.Job{}
	query := `SELECT id, payload, status, result, created_at, updated_at FROM jobs WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt)
	return job, err
}

func (r *jobRepo) Update(job *models.Job) error {
	job.UpdatedAt = time.Now()
	query := `UPDATE jobs SET status=?, result=?, updated_at=? WHERE id=?`
	_, err := r.db.Exec(query, job.Status, job.Result, job.UpdatedAt, job.ID)
	return err
}

func (r *jobRepo) List(offset, limit int) ([]models.Job, error) {
	query := `SELECT id, payload, status, result, created_at, updated_at FROM jobs ORDER BY id DESC LIMIT ? OFFSET ?`
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var job models.Job
		rows.Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt)
		jobs = append(jobs, job)
	}
	return jobs, nil
}
