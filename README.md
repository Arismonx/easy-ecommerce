# Easy E-Commerce Web Application

## Overview

This is a full-stack e-commerce web application built with modern web technologies, including Next.js, Go Fiber, GORM, PostgreSQL, and Docker. The project follows a microservice architecture with a focus on scalability, performance, and maintainability.

## Tech Stack

- **Frontend:** Next.js, React, TypeScript, pnpm
- **Backend:** Go (Fiber), GORM
- **Database:** PostgreSQL
- **Containerization:** Docker
- **Package Manager:** pnpm

## Features

- User authentication and authorization
- Product listing and search functionality
- Shopping cart and checkout system
- Order management for users and admins
- Payment integration (optional)
- Responsive design for mobile and desktop

## Installation & Setup

### Prerequisites

- Node.js & pnpm
- Go (latest version)
- Docker & Docker Compose
- PostgreSQL (if running locally without Docker)

### Clone the Repository

```sh
git clone https://github.com/Arismonx/easy-ecommerce.git
cd easy-ecommerce
```

### Setup Environment Variables

Create a `.env` file in the root directory and configure the necessary variables:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydatabase
```

### Running the Application

#### 1. Start Backend (Go Fiber)

```sh
cd backend
go mod tidy
go run main.go
```

#### 2. Start Frontend (Next.js)

```sh
cd frontend/my-app  
pnpm install
pnpm dev
```

#### 3. Run with Docker (Optional)

```sh
docker-compose up -d
```

## Project Structure

```
/ecommerce-project
│── backend/              # Go Fiber backend service
│── frontend/my-app       # Next.js frontend
│── .env                  # Environment variables
│── docker-compose.yml
│── README.md
```

## Contributing

Feel free to fork this repository, submit issues, and contribute to improving the project.

## License

This project is licensed under the MIT License.

