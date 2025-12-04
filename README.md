# Go Auth JWT Project

A robust backend service for User Authentication and Management implemented in Go (Golang) using JWT (JSON Web Tokens).

## 🚀 Features

- **User Registration**: Create new user accounts with secure password hashing (Bcrypt).
- **User Login**: Authenticate users and issue JWT tokens.
- **Protected Routes**: Middleware to verify JWT tokens for accessing protected endpoints.
- **User Profile**: Retrieve authenticated user details.
- **Clean Architecture**: Modular code structure (Handler, Service, Repository, Model).
- **Robust Configuration**: Environment variable management with fallback support.

## 🛠️ Tech Stack

- **Language**: Go (Golang) 1.22+
- **Database**: MySQL
- **Authentication**: JWT (JSON Web Tokens)
- **Routing**: Standard `net/http` ServeMux
- **Drivers & Libraries**:
  - `github.com/go-sql-driver/mysql` (MySQL Driver)
  - `github.com/golang-jwt/jwt/v5` (JWT)
  - `github.com/joho/godotenv` (Environment Variables)
  - `golang.org/x/crypto` (Password Hashing)

## ⚙️ Setup & Installation

### 1. Clone the Repository
```bash
git clone <repository-url>
cd auth-jwt-golang
```

### 2. Database Setup
Ensure you have MySQL installed and running. Create a database named `auth_db` (or whatever you prefer).

```sql
CREATE DATABASE auth_db;

USE auth_db;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 3. Environment Variables
Create a `.env` file in the root directory. You can copy the example below:

```env
APP_NAME=GoAuth
JWT_SECRET=your_super_secret_key_change_this
DB_HOST=127.0.0.1
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=auth_db
DB_PORT=3306
```

### 4. Install Dependencies
```bash
go mod tidy
```

## 🏃‍♂️ How to Run

You can run the application from the project root:

```bash
go run cmd/api/main.go
```

The server will start at `http://localhost:8080`.

## 📚 API Documentation

### Auth

#### Register
- **URL**: `/register`
- **Method**: `POST`
- **Body**:
  ```json
  {
      "name": "John Doe",
      "email": "john@example.com",
      "password": "password123"
  }
  ```

#### Login
- **URL**: `/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
      "email": "john@example.com",
      "password": "password123"
  }
  ```
- **Response**: Returns a JWT token.

### User

#### Get Profile (Protected)
- **URL**: `/users/profile`
- **Method**: `GET`
- **Headers**:
  - `Authorization`: `Bearer <your_token_here>`
- **Response**: Returns the authenticated user's profile.

## 📂 Project Structure

```
cmd/
└── api/
    └── main.go               # Entry point
internal/
├── auth/                     # Auth module (Login, Register)
├── config/                   # Configuration loader
├── database/                 # Database connection
├── middleware/               # Auth middleware
├── pkg/                      # Utilities (Hash, Response, Validator)
├── router/                   # Route definitions
├── user/                     # User module (Profile)
└── utils/                    # Helper functions
```
