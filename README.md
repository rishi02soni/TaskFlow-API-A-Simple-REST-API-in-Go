# 🚀 TaskFlow API – REST API in Go

TaskFlow API is a simple and clean RESTful API built using Go. It allows you to manage tasks with full CRUD operations.

---

##  Features

-  Create Task  
-  Get All Tasks  
-  Get Task by ID  
-  Update Task  
-  Delete Task  
-  Fast and lightweight (Go-powered)  
-  In-memory storage (no database required)

---

## 🛠️ Tech Stack

- Go (Golang)
- Gorilla Mux (Router)
- JSON (Data format)

---

##  Project Structure
```
taskflow-api/
│── main.go
│── go.mod
│── README.md
```


## ⚙️ Installation & Setup

### 1️ Clone the Repository

```
git clone https://github.com/yourusername/taskflow-api.git
cd taskflow-api
```
## Install Dependencies
```
go mod tidy
```
## Run the Server
```
go run main.go
```
Server will start at:
```
http://localhost:8080
```
## API Endpoints
```
{
  "title": "Master Go",
  "done": true
}
```
## Future Improvements
Add JWT Authentication
-Integrate Database (PostgreSQL / MongoDB)
-Docker Support
- Swagger API Documentation
- Deploy on Cloud (AWS / Azure / Render)

## Support

If you like this project, give it a ⭐ on GitHub!

