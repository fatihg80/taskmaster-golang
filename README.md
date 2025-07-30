# TaskMaster Go

![Go](https://img.shields.io/badge/Go-1.21+-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15+-blue)
![License](https://img.shields.io/badge/License-MIT-green)

A production-ready Kanban task management system built with Golang and PostgreSQL.

## Features

- ✅ Full CRUD operations for tasks
- 🗂️ Three workflow states (Todo/In Progress/Done)
- 🔄 RESTful API endpoints
- 🔒 JWT Authentication
- 📊 PostgreSQL database backend
- ✨ Live reload with Air
- 🧪 Unit tests with 85%+ coverage

## Quick Start

### Prerequisites
- Go 1.21+
- PostgreSQL 15+
- Air (for development)

```bash
# 1. Clone repository
git clone https://github.com/fatihg80/taskmaster-golang.git
cd taskmaster-golang

# 2. Setup database
createdb taskmaster
psql taskmaster -f migrations/001_init.sql

# 3. Configure environment
cp .env.example .env
nano .env  # Edit your DB credentials

JWT_SECRET_KEY=super_secret_stuff
POSTGRES_CONN_URL=postgresql://<YOUR_ACCOUNT>:<YOUR_API_KEY>@<YOUR_REGION>.sql.xata.sh/<YOUR_DATABASE_NAME>:main?sslmode=require

# 4. Install dependencies
go mod download

# 5. Run (development mode)
air

# 6. Run (production mode)
go build -o taskmaster
./taskmaster


#API Documentation :
Endpoint	Method	Description
_________________________________________
/api/v1/tasks	GET	List all tasks
_________________________________________
/api/v1/tasks	POST	Create new task
_________________________________________
/api/v1/tasks/{id}	GET	Get task details
_________________________________________
/api/v1/tasks/{id}	PUT	Update task
_________________________________________
/api/v1/tasks/{id}	DELETE	Remove task
_________________________________________



# Sample Task Object:
{
  "id": "uuidv4",
  "title": "Implement auth",
  "description": "Add JWT authentication",
  "status": "todo",
  "created_at": "2023-11-20T12:00:00Z"
}

# Development:

# Run tests
go test ./... -cover

# Generate swagger docs
swag init

# Build Docker image
docker build -t taskmaster .


# Deployment:
https://www.herokucdn.com/deploy/button.svg

# Docker compose
docker-compose up -d

# Kubernetes
kubectl apply -f k8s/


# Contributing:

    Fork the repository

    Create your feature branch (git checkout -b feature/amazing-feature)

    Commit your changes (git commit -m 'Add some amazing feature')

    Push to the branch (git push origin feature/amazing-feature)

    Open a Pull Request


# License:

Distributed under the MIT License. See LICENSE for more information.


**Key Features:**
- Ready-to-copy markdown format
- Badges for key technologies
- Clear installation steps
- API documentation table
- Deployment options
- Contribution guidelines
- Mobile-responsive layout

**To Use:**
1. Copy this entire content
2. Create `README.md` in your project root
3. Paste and customize the sections as needed
4. Update the API endpoints to match your actual implementation

Would you like me to add any specific sections like:
- Architecture diagram
- Screenshots
- Detailed API examples
- Roadmap
- Troubleshooting guide?



