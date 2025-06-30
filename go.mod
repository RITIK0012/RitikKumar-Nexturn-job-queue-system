module job-queue-system

go 1.23.4

require github.com/go-sql-driver/mysql v1.9.3

require go.uber.org/multierr v1.10.0 // indirect

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	go.uber.org/zap v1.27.0
)
