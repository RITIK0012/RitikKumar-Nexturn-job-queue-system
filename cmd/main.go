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
	config.InitSchema() // <-- Added this line to auto-create the jobs table
	logger := config.NewLogger()

	repo := repository.NewJobRepository(db)
	service := services.NewJobService(repo, logger)
	handler := handlers.NewJobHandler(service)

	queue.StartWorkerPool(service.ProcessJob, 5)

	http.HandleFunc("/job", handler.HandleJobSubmission)
	http.HandleFunc("/job/", handler.HandleJobStatus)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
