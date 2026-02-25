# Local REST API Server

Complete REST API server with all HTTP methods (GET, POST, PUT, PATCH, DELETE) and Bearer token authentication.

## Quick Start

```bash
cd server
go run server.go
```

Server starts on `http://localhost:8080`

Valid Token: `secret_token_12345`

## API Endpoints

### Health Check (No Auth Required)

```http
GET http://localhost:8080/health
```

Response:
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

---

### User Endpoints (Auth Required)

#### Get User
```http
GET http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
```

Response:
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "username": "johndoe"
}
```

#### Create User
```http
POST http://localhost:8080/users
Authorization: Bearer secret_token_12345
Content-Type: application/json

{
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "username": "janesmith"
}
```

Response (201 Created):
```json
{
  "id": 2,
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "username": "janesmith"
}
```

#### Update User (Full - PUT)
```http
PUT http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
Content-Type: application/json

{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "username": "johnupdated"
}
```

Response:
```json
{
  "id": 1,
  "name": "John Updated",
  "email": "john.updated@example.com",
  "username": "johnupdated"
}
```

#### Update User (Partial - PATCH)
```http
PATCH http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
Content-Type: application/json

{
  "name": "John Patched"
}
```

Response:
```json
{
  "id": 1,
  "name": "John Patched",
  "email": "john.doe@example.com",
  "username": "johndoe"
}
```

#### Delete User
```http
DELETE http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
```

Response:
```json
{
  "message": "User deleted successfully",
  "data": {
    "id": 1
  }
}
```

---

### Post Endpoints (Auth Required)

#### Get Post
```http
GET http://localhost:8080/posts/1
Authorization: Bearer secret_token_12345
```

Response:
```json
{
  "userId": 1,
  "id": 1,
  "title": "My First Post",
  "body": "This is the content of my first post..."
}
```

#### Create Post
```http
POST http://localhost:8080/posts
Authorization: Bearer secret_token_12345
Content-Type: application/json

{
  "userId": 1,
  "title": "New Post Title",
  "body": "This is the body of the new post"
}
```

Response (201 Created):
```json
{
  "userId": 1,
  "id": 2,
  "title": "New Post Title",
  "body": "This is the body of the new post"
}
```

#### Update Post (PUT)
```http
PUT http://localhost:8080/posts/1
Authorization: Bearer secret_token_12345
Content-Type: application/json

{
  "userId": 1,
  "title": "Updated Post Title",
  "body": "This post has been updated"
}
```

Response:
```json
{
  "userId": 1,
  "id": 1,
  "title": "Updated Post Title",
  "body": "This post has been updated"
}
```

#### Delete Post
```http
DELETE http://localhost:8080/posts/1
Authorization: Bearer secret_token_12345
```

Response:
```json
{
  "message": "Post deleted successfully",
  "data": {
    "id": 1
  }
}
```

---

## Testing with Postman

### Collection Setup

Create a Postman collection with these requests:

1. **Environment Variables**
   - `base_url`: `http://localhost:8080`
   - `token`: `secret_token_12345`

2. **Authorization Setup**
   - Type: Bearer Token
   - Token: `{{token}}`

### Test Scenarios

#### Scenario 1: CRUD User Flow
1. GET /users/1 (read existing)
2. POST /users (create new)
3. PUT /users/1 (full update)
4. PATCH /users/1 (partial update)
5. DELETE /users/1 (delete)

#### Scenario 2: CRUD Post Flow
1. GET /posts/1 (read existing)
2. POST /posts (create new)
3. PUT /posts/1 (update)
4. DELETE /posts/1 (delete)

#### Scenario 3: Authentication Tests
1. GET /health (no auth - should work)
2. GET /users/1 without token (should fail 401)
3. GET /users/1 with wrong token (should fail 401)
4. GET /users/1 with valid token (should work)

---

## Testing with curl

### GET Request
```bash
curl -H "Authorization: Bearer secret_token_12345" \
     http://localhost:8080/users/1
```

### POST Request
```bash
curl -X POST \
     -H "Authorization: Bearer secret_token_12345" \
     -H "Content-Type: application/json" \
     -d '{"name":"Jane Smith","email":"jane@example.com","username":"janesmith"}' \
     http://localhost:8080/users
```

### PUT Request
```bash
curl -X PUT \
     -H "Authorization: Bearer secret_token_12345" \
     -H "Content-Type: application/json" \
     -d '{"name":"John Updated","email":"john.updated@example.com","username":"johnupdated"}' \
     http://localhost:8080/users/1
```

### PATCH Request
```bash
curl -X PATCH \
     -H "Authorization: Bearer secret_token_12345" \
     -H "Content-Type: application/json" \
     -d '{"name":"John Patched"}' \
     http://localhost:8080/users/1
```

### DELETE Request
```bash
curl -X DELETE \
     -H "Authorization: Bearer secret_token_12345" \
     http://localhost:8080/users/1
```

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

### 401 Unauthorized - Missing Token
```json
{
  "error": "Missing Authorization header"
}
```

### 401 Unauthorized - Invalid Token
```json
{
  "error": "Invalid token"
}
```

### 405 Method Not Allowed
```json
{
  "error": "Method not allowed"
}
```

---

## HTTP Methods Explained

| Method | Purpose | Idempotent | Request Body | Response Body |
|--------|---------|------------|--------------|---------------|
| GET | Retrieve resource | Yes | No | Yes |
| POST | Create resource | No | Yes | Yes |
| PUT | Replace resource | Yes | Yes | Yes |
| PATCH | Update resource partially | No | Yes | Yes |
| DELETE | Remove resource | Yes | No | Yes |

### PUT vs PATCH

- **PUT**: Replaces the entire resource (all fields required)
- **PATCH**: Updates only specified fields (partial update)

Example:
```json
// PUT - Must send all fields
{
  "name": "John",
  "email": "john@example.com",
  "username": "john"
}

// PATCH - Send only what you want to change
{
  "name": "John Updated"
}
```

---

## Testing with Go Client

The Go client in the main project automatically tests all endpoints:

```bash
# Terminal 1 - Start server
cd server
go run server.go

# Terminal 2 - Run client tests
cd ..
go run .
```

You'll see all HTTP methods tested automatically!

---

## Extending the Server

### Add New Endpoint

```go
http.HandleFunc("/custom", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        // Handle GET
    case http.MethodPost:
        // Handle POST
    default:
        respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
    }
}))
```

### Add Request Logging

```go
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next(w, r)
    }
}
```

### Add CORS Support

```go
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
```
