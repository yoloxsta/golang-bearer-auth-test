# Postman Testing Guide

## 🚀 Quick Start

### Prerequisites
- Server is running on `http://localhost:8080`
- Postman is installed

### Import Collection

1. Open Postman
2. Click "Import" button (top left)
3. Click "Upload Files"
4. Select `postman_collection.json` from this project
5. Click "Import"

You'll see "Go REST API Lab" collection with all requests ready!

---
##  Test Scenarios

### Scenario 1: Health Check (No Auth Required)

**Request:**
```
GET http://localhost:8080/health
```

**Steps:**
1. Open "Health Check" request
2. Click "Send"

**Expected Response (200 OK):**
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

✅ This proves the server is running!

---

### Scenario 2: GET User (With Auth)

**Request:**
```
GET http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
```

**Steps:**
1. Open "Users" folder → "Get User"
2. Check "Authorization" tab shows Bearer token
3. Click "Send"

**Expected Response (200 OK):**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "username": "johndoe"
}
```

✅ Bearer token authentication works!

---

### Scenario 3: POST Create User

**Request:**
```
POST http://localhost:8080/users
Authorization: Bearer secret_token_12345
Content-Type: application/json

Body:
{
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "username": "janesmith"
}
```

**Steps:**
1. Open "Users" folder → "Create User"
2. Check "Body" tab shows JSON data
3. Click "Send"

**Expected Response (201 Created):**
```json
{
  "id": 2,
  "name": "Jane Smith",
  "email": "jane.smith@example.com",
  "username": "janesmith"
}
```

✅ POST request creates new resource!

---

### Scenario 4: PUT Update User (Full)

**Request:**
```
PUT http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
Content-Type: application/json

Body:
{
  "name": "John Updated",
  "email": "john.updated@example.com",
  "username": "johnupdated"
}
```

**Steps:**
1. Open "Users" folder → "Update User (PUT)"
2. Check all fields in Body
3. Click "Send"

**Expected Response (200 OK):**
```json
{
  "id": 1,
  "name": "John Updated",
  "email": "john.updated@example.com",
  "username": "johnupdated"
}
```

✅ PUT replaces entire resource!

---

### Scenario 5: PATCH Update User (Partial)

**Request:**
```
PATCH http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
Content-Type: application/json

Body:
{
  "name": "John Patched"
}
```

**Steps:**
1. Open "Users" folder → "Update User (PATCH)"
2. Notice only "name" field in Body
3. Click "Send"

**Expected Response (200 OK):**
```json
{
  "id": 1,
  "name": "John Patched",
  "email": "john.doe@example.com",
  "username": "johndoe"
}
```

✅ PATCH updates only specified fields!  
Notice: email and username unchanged!

---

### Scenario 6: DELETE User

**Request:**
```
DELETE http://localhost:8080/users/1
Authorization: Bearer secret_token_12345
```

**Steps:**
1. Open "Users" folder → "Delete User"
2. Click "Send"

**Expected Response (200 OK):**
```json
{
  "message": "User deleted successfully",
  "data": {
    "id": 1
  }
}
```

✅ DELETE removes resource!

---

### Scenario 7: Test Posts (Same Pattern)

Repeat the same tests for Posts:

1. **GET Post** - `GET /posts/1`
2. **Create Post** - `POST /posts`
3. **Update Post** - `PUT /posts/1`
4. **Delete Post** - `DELETE /posts/1`

All work the same way as Users!

---

## 🔐 Authentication Tests

### Test 1: No Token (Should Fail)

**Steps:**
1. Open "Auth Tests" folder → "No Token (Should Fail)"
2. Notice no Authorization header
3. Click "Send"

**Expected Response (401 Unauthorized):**
```json
{
  "error": "Missing Authorization header"
}
```

✅ Server rejects requests without token!

---

### Test 2: Wrong Token (Should Fail)

**Steps:**
1. Open "Auth Tests" folder → "Wrong Token (Should Fail)"
2. Check Authorization shows "wrong_token"
3. Click "Send"

**Expected Response (401 Unauthorized):**
```json
{
  "error": "Invalid token"
}
```

✅ Server validates token correctly!

---

## 📊 Understanding Responses

### Status Codes

| Code | Meaning | When You See It |
|------|---------|-----------------|
| 200 | OK | Successful GET, PUT, PATCH, DELETE |
| 201 | Created | Successful POST |
| 400 | Bad Request | Invalid JSON or missing fields |
| 401 | Unauthorized | Missing or invalid token |
| 405 | Method Not Allowed | Wrong HTTP method |

### Response Headers

Check the "Headers" tab in response to see:
- `Content-Type: application/json`
- Status code and reason

---

## 🎯 Complete Test Flow

Run these in order to test everything:

1. ✅ Health Check (no auth)
2. ✅ GET User (with auth)
3. ✅ POST User (create)
4. ✅ PUT User (full update)
5. ✅ PATCH User (partial update)
6. ✅ DELETE User
7. ✅ GET Post
8. ✅ POST Post
9. ✅ PUT Post
10. ✅ DELETE Post
11. ✅ No Token Test (should fail)
12. ✅ Wrong Token Test (should fail)

---

## 🔧 Customizing Requests

### Change the Token

1. Click collection name "Go REST API Lab"
2. Go to "Variables" tab
3. Change `token` value
4. Save

All requests will use the new token!

### Change the URL

1. Click collection name
2. Go to "Variables" tab
3. Change `base_url` value
4. Save

Now you can test against different servers!

### Modify Request Body

1. Open any POST/PUT/PATCH request
2. Go to "Body" tab
3. Edit the JSON
4. Click "Send"

Try different data!

---

## 💡 Tips & Tricks

### Save Responses

Click "Save Response" to keep examples for later.

### Use Tests Tab

Add automatic tests in the "Tests" tab:

```javascript
pm.test("Status is 200", function () {
    pm.response.to.have.status(200);
});

pm.test("Response has id", function () {
    var jsonData = pm.response.json();
    pm.expect(jsonData).to.have.property('id');
});
```

### Use Pre-request Scripts

Add dynamic data in "Pre-request Script":

```javascript
pm.variables.set("timestamp", Date.now());
```

### Run Collection

Click "Run" button to execute all requests automatically!

---

## 🐛 Troubleshooting

### "Could not get response"
- Check server is running
- Verify URL is `http://localhost:8080`
- Check firewall settings

### "401 Unauthorized"
- Verify token is `secret_token_12345`
- Check Authorization type is "Bearer Token"
- Make sure token is in collection variables

### "400 Bad Request"
- Check JSON syntax in Body
- Verify all required fields are present
- Check Content-Type is `application/json`

### "Connection refused"
- Server is not running
- Start server: `cd server && go run server.go`

---

## 📸 Visual Guide

### Setting Up Authorization

1. Click request
2. Go to "Authorization" tab
3. Type: Select "Bearer Token"
4. Token: Enter `{{token}}` (uses collection variable)

### Editing Request Body

1. Click request
2. Go to "Body" tab
3. Select "raw"
4. Select "JSON" from dropdown
5. Edit JSON content
6. Click "Send"

### Viewing Response

After clicking "Send":
- **Body** tab: See JSON response
- **Headers** tab: See response headers
- **Status**: See status code (200, 201, 401, etc.)
- **Time**: See response time

---

## ✅ Success Checklist

After testing, you should have:

- [x] Health check works without auth
- [x] GET requests return data
- [x] POST requests create resources
- [x] PUT requests update all fields
- [x] PATCH requests update partial fields
- [x] DELETE requests remove resources
- [x] Invalid tokens are rejected
- [x] Missing tokens are rejected
- [x] All status codes are correct
- [x] All responses are JSON

---

## 🎉 You're Done!

You've successfully tested:
- ✅ All HTTP methods (GET, POST, PUT, PATCH, DELETE)
- ✅ Bearer token authentication
- ✅ Request/response handling
- ✅ Error scenarios
- ✅ Complete REST API flow

Your API is working perfectly! 🚀

---

## 📚 Next Steps

1. Try modifying request bodies
2. Test with invalid data
3. Create your own requests
4. Export collection to share
5. Add automated tests
6. Integrate with your own APIs

Happy testing! 🎯
