package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"job-queue-system/models"
)

type JobRepository interface {
	Create(job *models.Job) error
	FindByID(id int64) (*models.Job, error)
	Update(job *models.Job) error
	List(offset, limit int) ([]models.Job, error)
}

type jobRepo struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) JobRepository {
	return &jobRepo{db}
}

func (r *jobRepo) Create(job *models.Job) error {
	now := time.Now()
	job.CreatedAt = now
	job.UpdatedAt = now

	query := `INSERT INTO jobs (payload, status, created_at, updated_at)
	          VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRow(query, job.Payload, job.Status, job.CreatedAt, job.UpdatedAt).Scan(&job.ID)
	return err
}

func (r *jobRepo) FindByID(id int64) (*models.Job, error) {
	job := &models.Job{}
	query := `SELECT id, payload, status, result, created_at, updated_at FROM jobs WHERE id = $1`
	err := r.db.Get(job, query, id)
	return job, err
}

func (r *jobRepo) Update(job *models.Job) error {
	job.UpdatedAt = time.Now()
	query := `UPDATE jobs SET status=$1, result=$2, updated_at=$3 WHERE id=$4`
	_, err := r.db.Exec(query, job.Status, job.Result, job.UpdatedAt, job.ID)
	return err
}

func (r *jobRepo) List(offset, limit int) ([]models.Job, error) {
	var jobs []models.Job
	query := `SELECT id, payload, status, result, created_at, updated_at
	          FROM jobs ORDER BY id DESC LIMIT $1 OFFSET $2`
	err := r.db.Select(&jobs, query, limit, offset)
	return jobs, err
}
