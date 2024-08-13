# API Security

A simple Go API with JWT authentication, rate limiting, and real-time endpoints. This API demonstrates basic security practices and provides two functional endpoints to check the rate-limiting: `/current-time` and `/echo`.

## Project Structure

- **`/cmd`**: Contains the entry point of the application.
- **`/internal`**: Contains internal application logic including:
  - **`/config`**: Configuration loading (e.g., from `.env` files).
  - **`/handlers`**: HTTP request handlers.
  - **`/middleware`**: Middlewares for authentication and rate limiting.
  - **`/services`**: Utility functions, such as JWT handling.
- **`/pkg`**: Contains packages used across the application, such as Redis client setup.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Redis server (can be run locally or use a hosted service)
- Docker (optional, for running Redis in a container)

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/balasl342/api-security.git
   cd api-security

2. **Install Dependencies:**

    ```sh
    go mod tidy

3. **Create a .env file in the root directory and add your environment variables:**

    ```sh
    SECRET_KEY=your_secret_key

4. **Set up Redis:**
    Note : Ensure Redis is running locally on the default port (6379) or use the docker.
    ```sh
    docker run -d --name redis -p 6379:6379 redis

### Running the Application
- go run cmd/main.go
