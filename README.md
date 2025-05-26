# 📚 Book Tracker API

A RESTful API built with Go and Fiber for managing a personal book collection. It supports user authentication via JWT, allowing users to register, log in, and perform CRUD operations on their books. The application uses SQLite for data storage and is containerized with Docker for easy deployment.

---

## 🚀 Features

* **User Authentication**: Secure registration and login using JWT.
* **Book Management**: Create, read, update, and delete books.
* **SQLite Integration**: Lightweight database for storing user and book data.
* **Dockerized**: Easily deployable using Docker.
* **Modular Structure**: Organized codebase with clear separation of concerns.

---

## 🛠️ Technologies Used

* **Go**: Programming language.
* **Fiber**: Web framework for Go.
* **GORM**: ORM library for Go.
* **SQLite**: Relational database.
* **JWT**: JSON Web Tokens for authentication.
* **Docker**: Containerization platform.

---

## 📦 Installation

### Prerequisites

* [Go](https://golang.org/doc/install) installed.
* [Docker](https://docs.docker.com/get-docker/) installed (optional, for containerization).

### Clone the Repository

```bash
git clone https://github.com/EdsonV1/book-tracker-api.git
cd book-tracker-api
```

### Set Environment Variables

Create a .env file in the root directory with the following content:

```bash
JWT_SECRET=your_jwt_secret_key
```

You could use a JWT generator:

* [JWT Secret Generator](https://jwtsecret.com/generate)

---

## How to run it?

### Build and Run with Go

```bash
go mod tidy
go run main.go
```

### Build and Run with Docker

```bash
docker build -t book-tracker-api .
docker run -p 3000:3000 book-tracker-api
```

---

## 📚 API Endpoints

> All responses are in JSON format.
> Endpoints related to books require a valid JWT token in the `Authorization` header.

### 🔐 Authentication

#### `POST /auth/register`

Registers a new user.

**Request Body:**

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

#### `POST /auth/login`

Authenticates a user and returns a JWT token.

**Request Body:**

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Response:**

```json
{
  "token": "your_jwt_token"
}
```

---

### 📖 Books


#### `GET /books`

Fetch all books.

**Response:**

```json
[
  {
    "id": 1,
    "title": "Clean Code",
    "author": "Robert C. Martin",
    "read": true
  }
]
```

#### `GET /books/:id`

Fetch a specific book by its ID.

**Response:**

```json
{
  "id": 1,
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "read": true
}
```

> ⚠️ All routes below require a valid JWT in the request header:
>
> ```
> Authorization: Bearer your_jwt_token
> ```

#### `POST /books`

Create a new book.

**Request Body:**

```json
{
  "title": "The Pragmatic Programmer",
  "author": "Andrew Hunt",
  "read": false
}
```

#### `PUT /books/:id`

Update an existing book by its ID.

**Request Body:**

```json
{
  "title": "The Pragmatic Programmer",
  "author": "Andrew Hunt",
  "read": true
}
```

#### `DELETE /books/:id`

Delete a book by its ID.

**Response:**

```json
{
  "message": "Book deleted"
}
```
