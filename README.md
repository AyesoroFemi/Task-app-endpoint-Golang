# Task_app
Fullstack Web Task listing App
Task management app in golang using chi router
- env: contains the connection to the mongodb used for connection
- go mod init github.com/todo_app
- go mod tidy to install all other packages
- Install react js using vite : npm create vite@latest
- Backend Endpoint
- To get list of tasks:  http://localhost:8080/api/v1/todos
- To get single tasks: http://localhost:8080/api/v1/todos/{id}
- To create task: http://localhost:8080/api/v1/todos/create
- To update task: http://localhost:8080/api/v1/todos/update/{id}
- To delete task: http://localhost:8080/api/v1/todos/delete/{id}

TECH STACKS
-BACKEND: Golang, Chi Router
-FRONTEND: React js