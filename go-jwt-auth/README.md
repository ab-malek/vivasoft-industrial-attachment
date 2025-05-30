
#  🛡️ JWT Authentication Web App using Go (Gin Framework)



This is a simple web application that implements JWT (JSON Web Token) authentication using the Go programming language and the Gin web framework. It supports secure user login, protected routes, and middleware-based token validation.

---

## 🚀 Features

- User registration and login
- JWT token generation on successful login
- Secure cookie-based token storage
- Protected routes using middleware (`RequireAuth`)
- Token validation with expiration checks
- Integration with a PostgreSQL/MySQL/SQLite database via GORM (Go ORM)


## Authors

- [@ab-malek](https://www.github.com/ab-malek)


# API Endpoints

### Public
POST /signup  - Register a new user
POST /login - Log in and receive JWT in cookie

### Protected
GET /login - Returns authenticated user's profile

# Dependencies
Gin: go get -u github.com/gin-gonic/gin

GORM: go get -u gorm.io/gorm

mysql Driver: go get -u gorm.io/driver/mysql/

JWT: go get github.com/golang-jwt/jwt/v5

Godotenv: go get github.com/joho/godotenv
