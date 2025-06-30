package main

import (
	"log"
	"net/http"

	"job-queue-system/config"
	"job-queue-system/handlers"
	"job-queue-system/queue"
	"job-queue-system/repository"
	"job-queue-system/services"
)

func main() {
	db := config.ConnectDB()
	logger := config.NewLogger()

	repo := repository.NewJobRepository(db)
	service := services.NewJobService(repo, logger)
	handler := handlers.NewJobHandler(service)

	// Start workers using the service.ProcessJob callback
	queue.StartWorkerPool(service.ProcessJob, 5)

	// ✅ Fixed paths
	http.HandleFunc("/job", handler.HandleJobSubmission)   // POST
	http.HandleFunc("/job/", handler.HandleJobStatus)      // GET /job/{id}

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
