#  Go Job Queue System (MySQL + Docker)

This is a simple job queue backend built in Go. It uses:

-  MySQL for database
-  Docker to run MySQL easily
-  Goroutines + channel for background processing
-  REST API (POST for submitting job, GET for checking status)

---

##  Basic Idea

1. You send a job (some data) using POST `/job`.
2. The job is saved to MySQL with status = `queued`.
3. A worker picks the job and processes it.
4. Once done, it updates the job status to `completed`.
5. You can check the result using GET `/job/{id}`.

---

##  Step-by-Step Guide

### 1.  Start MySQL using Docker

Open PowerShell and run this command to start a MySQL container:

```powershell
docker run -d --name docker-mysql -p 3306:3306 ^
  -e MYSQL_ROOT_PASSWORD=root ^
  -e MYSQL_DATABASE=jobqueue ^
  -e MYSQL_USER=myuser ^
  -e MYSQL_PASSWORD=mypassword ^
  mysql:latest
This creates a MySQL container with:

Database: jobqueue

Username: myuser

Password: mypassword

2.  Set Environment Variable in PowerShell
Tell your Go app how to connect to MySQL:

powershell
Copy code
$env:DB_DSN="myuser:mypassword@tcp(127.0.0.1:3306)/jobqueue"
3. 🏃 Run Your Go Application
Go to your project folder and start the server:

powershell
Copy code
cd "C:\Users\HP\Desktop\job-queue-system"
go run cmd/main.go
You should see:

arduino
Copy code
Server running on port 8080
4. 🧪 Test API in Postman
 Submit a Job
Method: POST

URL: http://localhost:8080/job

Body (raw JSON):

json
Copy code
{
  "payload": "hello world"
}
You’ll get a response like:

json
Copy code
{
  "id": 1,
  "payload": "hello world",
  "status": "queued",
  "result": "",
  "created_at": "...",
  "updated_at": "..."
}
 Check Job Status
Method: GET

URL: http://localhost:8080/job/1

Response after a few seconds:

json
Copy code
{
  "id": 1,
  "payload": "hello world",
  "status": "completed",
  "result": "Processed: hello world",
  "created_at": "...",
  "updated_at": "..."
}
 SQL Table You Need in MySQL
Before testing, create this table in the jobqueue database using DBeaver or any MySQL client:

sql
Copy code
CREATE TABLE jobs (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  payload TEXT NOT NULL,
  status VARCHAR(50) NOT NULL,
  result TEXT,
  created_at DATETIME,
  updated_at DATETIME
);
 API Summary
Method	Endpoint	Description
POST	/job	Submit new job
GET	/job/{id}	Get job status#   R i t i k K u m a r - N e x t u r n - j o b - q u e u e - s y s t e m  
 