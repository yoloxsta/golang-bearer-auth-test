# Testing Guide

Complete guide for testing all REST API methods with Bearer token authentication.

## Quick Test

### 1. Start Server
```bash
cd server
go run server.go
```

### 2. Run Client Tests
Open new terminal:
```bash
copy .env.example .env
go run .
```

You'll see all HTTP methods tested automatically! ✅

---

## Manual Testing with Postman

### Setup

1. **Create Environment**
   - Name: `Local API`
   - Variables:
     - `base_url`: `http://localhost:8080`
     - `token`: `secret_token_12345`

2. **Set Authorization**
   - Type: `Bearer Token`
   - Token: `{{token}}`

### Test Collection

#### 1. Health Check (No Auth)
```
GET {{base_url}}/health
```
Expected: 200 OK

#### 2. GET User
```
GET {{base_url}}/users/1
Authorization: Bearer {{token}}
```
Expected: 200 OK with user data

#### 3. POST User
```
POST {{base_url}}/users
Authorization: Bearer {{token}}
Content-Type: application/json

Body:
{
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "username": "janesmith"
}
```
Expected: 201 Created

#### 4. PUT User (Full Update)
```
PUT {{base_url}}/users/1
Authorization: Bearer {{token}}
Content-Type: application/json

Body:
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "username": "johnupdated"
}
```
Expected: 200 OK

#### 5. PATCH User (Partial Update)
```
PATCH {{base_url}}/users/1
Authorization: Bearer {{token}}
Content-Type: application/json

Body:
{
  "name": "John Patched"
}
```
Expected: 200 OK (only name changed)

#### 6. DELETE User
```
DELETE {{base_url}}/users/1
Authorization: Bearer {{token}}
```
Expected: 200 OK

#### 7. GET Post
```
GET {{base_url}}/posts/1
Authorization: Bearer {{token}}
```
Expected: 200 OK

#### 8. POST Post
```
POST {{base_url}}/posts
Authorization: Bearer {{token}}
Content-Type: application/json

Body:
{
  "userId": 1,
  "title": "My New Post",
  "body": "This is the content of my post"
}
```
Expected: 201 Created

#### 9. PUT Post
```
PUT {{base_url}}/posts/1
Authorization: Bearer {{token}}
Content-Type: application/json

Body:
{
  "userId": 1,
  "title": "Updated Post",
  "body": "Updated content"
}
```
Expected: 200 OK

#### 10. DELETE Post
```
DELETE {{base_url}}/posts/1
Authorization: Bearer {{token}}
```
Expected: 200 OK

---

## Authentication Tests

### Test 1: Missing Token
```
GET http://localhost:8080/users/1
(No Authorization header)
```
Expected: 401 Unauthorized
```json
{
  "error": "Missing Authorization header"
}
```

### Test 2: Invalid Token Format
```
GET http://localhost:8080/users/1
Authorization: secret_token_12345
```
Expected: 401 Unauthorized
```json
{
  "error": "Invalid Authorization format. Use: Bearer <token>"
}
```

### Test 3: Wrong Token
```
GET http://localhost:8080/users/1
Authorization: Bearer wrong_token
```
Expected: 401 Unauthorized
```json
{
  "error": "Invalid token"
}
```

### Test 4: Valid Token
```
GET http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
```
Expected: 200 OK with data

---

## Testing with curl

### Windows PowerShell

```powershell
# GET
curl -H "Authorization: Bearer secret_token_12345" http://localhost:8080/users/1

# POST
curl -X POST -H "Authorization: Bearer secret_token_12345" -H "Content-Type: application/json" -d '{\"name\":\"Jane\",\"email\":\"jane@example.com\",\"username\":\"jane\"}' http://localhost:8080/users

# PUT
curl -X PUT -H "Authorization: Bearer secret_token_12345" -H "Content-Type: application/json" -d '{\"name\":\"John Updated\",\"email\":\"john@example.com\",\"username\":\"john\"}' http://localhost:8080/users/1

# PATCH
curl -X PATCH -H "Authorization: Bearer secret_token_12345" -H "Content-Type: application/json" -d '{\"name\":\"John Patched\"}' http://localhost:8080/users/1

# DELETE
curl -X DELETE -H "Authorization: Bearer secret_token_12345" http://localhost:8080/users/1
```

---

## HTTP Methods Summary

| Method | Endpoint | Purpose | Auth | Body | Response |
|--------|----------|---------|------|------|----------|
| GET | /health | Health check | ❌ | ❌ | Status |
| GET | /users/1 | Get user | ✅ | ❌ | User object |
| POST | /users | Create user | ✅ | ✅ | New user |
| PUT | /users/1 | Update user (full) | ✅ | ✅ | Updated user |
| PATCH | /users/1 | Update user (partial) | ✅ | ✅ | Updated user |
| DELETE | /users/1 | Delete user | ✅ | ❌ | Success message |
| GET | /posts/1 | Get post | ✅ | ❌ | Post object |
| POST | /posts | Create post | ✅ | ✅ | New post |
| PUT | /posts/1 | Update post | ✅ | ✅ | Updated post |
| DELETE | /posts/1 | Delete post | ✅ | ❌ | Success message |

---

## Expected Responses

### Success Responses

#### GET User (200 OK)
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "username": "johndoe"
}
```

#### POST User (201 Created)
```json
{
  "id": 2,
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "username": "janesmith"
}
```

#### DELETE User (200 OK)
```json
{
  "message": "User deleted successfully",
  "data": {
    "id": 1
  }
}
```

### Error Responses

#### 400 Bad Request
```json
{
  "error": "Invalid request body"
}
```

#### 401 Unauthorized
```json
{
  "error": "Invalid token"
}
```

#### 405 Method Not Allowed
```json
{
  "error": "Method not allowed"
}
```

---

## Testing Checklist

### Basic Functionality
- [ ] Server starts without errors
- [ ] Health endpoint works without auth
- [ ] All endpoints require Bearer token
- [ ] Invalid token returns 401
- [ ] Missing token returns 401

### GET Requests
- [ ] GET /users/1 returns user data
- [ ] GET /posts/1 returns post data

### POST Requests
- [ ] POST /users creates new user
- [ ] POST /posts creates new post
- [ ] Missing required fields returns 400

### PUT Requests
- [ ] PUT /users/1 updates all fields
- [ ] PUT /posts/1 updates post

### PATCH Requests
- [ ] PATCH /users/1 updates only specified fields
- [ ] Unspecified fields remain unchanged

### DELETE Requests
- [ ] DELETE /users/1 returns success
- [ ] DELETE /posts/1 returns success

### Client Tests
- [ ] Go client connects successfully
- [ ] All HTTP methods work
- [ ] Errors are handled gracefully
- [ ] Context timeout works

---

## Troubleshooting

### Server won't start
```
Error: listen tcp :8080: bind: address already in use
```
Solution: Port 8080 is in use. Kill the process or change PORT in server.go

### Client can't connect
```
Error: connection refused
```
Solution: Make sure server is running first

### 401 Unauthorized
```
Error: Invalid token
```
Solution: Check .env file has correct token: `secret_token_12345`

### Missing .env file
```
Error: failed to open .env file
```
Solution: Run `copy .env.example .env`

---

## Performance Testing

### Load Test with curl
```bash
# Run 100 requests
for i in {1..100}; do
  curl -H "Authorization: Bearer secret_token_12345" http://localhost:8080/users/1
done
```

### Concurrent Requests
```bash
# Run 10 concurrent requests
for i in {1..10}; do
  curl -H "Authorization: Bearer secret_token_12345" http://localhost:8080/users/1 &
done
wait
```

---

## Next Steps

1. ✅ Test all endpoints manually in Postman
2. ✅ Verify authentication works correctly
3. ✅ Test error scenarios
4. ✅ Run automated client tests
5. 📝 Document any issues found
6. 🚀 Ready to push to GitHub!

---

## Quick Reference

**Server**: `cd server && go run server.go`  
**Client**: `go run .`  
**Demo**: `go run . demo`  
**Token**: `secret_token_12345`  
**Base URL**: `http://localhost:8080`
