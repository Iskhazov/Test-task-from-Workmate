# Test-task-from-Workmate
A simple Golang-based API to demonstrate asynchronous task processing.
Requests are added via REST API and processed in the background.
Data is stored in a MySQL database with built-in migration support.
## Setup
1. Clone repository
```sh
git clone https://github.com/Iskhazov/Test-task-from-Workmate.git
cd awesomeProject
```
2. Install Dependencies
 ```sh
go mod tidy
```
3. Database Configuration  
* Set up Mysql database.  
* Configure connection details in config/env.go  
4. Run Migrations
 ```sh
make migrate-up
```
5. Start application
 ```sh
make run
```
## API Endpoints
POST /api/v1/requests - Make new request.  
GET /api/v1/requests - Get all requests.  
GET /api/v1/requests/{taskID} - GET task status by task ID.  

