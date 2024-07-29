
# Go Todo API Project

## Description

This project is a backend REST API service written in Go. It includes database setup and basic CRUD operations.

## Prerequisites

-   Go
-   PostgreSQL

## Getting Started

### Clone the Repository

`git clone git@github.com:illiyyin/sagala-todo.git`

`cd sagala-todo` 

### Set Up the Database

1.  **Install PostgreSQL:**
    
    Follow the instructions to install PostgreSQL on your system: [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
    
2.  **Create a Database:**
    
    `psql -U postgres` 
    
    `CREATE DATABASE your_database_name;
    CREATE USER your_username WITH PASSWORD 'your_password';
    GRANT ALL PRIVILEGES ON DATABASE your_database_name TO your_username;` 

3. **Copy Env**

    
    
    Create .env file based on .env.example and filled with data based on database you just made before. Or if you want to use cloud database, its same

   if this is your first time run this project, you can set `INIT_DB=1`

   
    	INIT_DB=1
    	DB_HOST=localhost
    	DB_USER=your_username
    	DB_PASSWORD=your_password
    	DB_NAME=your_database_name
	
    
5.  **Run the project:**
    
	### Install Dependencies
	`go mod download` 

	### Run the Project
	`go run main.go` 

### API Endpoints

after run the project, you can access the swagger here [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)

-   **GET /tasks**				: Get all tasks
-   **POST /task**				: Create a new task
-   **GET /task/[id]**		: Get a task by ID
-   **PATCH /task/[id]**	: Update a task by ID
-   **DELETE /task/[id]**	: Delete a task by ID

