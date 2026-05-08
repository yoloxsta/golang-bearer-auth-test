# Bearer Token Authentication - Code Explanation

## File 1: `.env` (Token Storage)

**Lines 2-3:**
```env
API_BASE_URL=http://localhost:8080
BEARER_TOKEN=secret_token_12345
```
- This is where the token is stored
- Change `secret_token_12345` to any token you want

---
## File 2: `server/server.go` (Backend - Go API)

### Part 1: Load Token from Environment

**Lines 12-14:**
```go
var (
	PORT        = "8080"
	VALID_TOKEN = os.Getenv("BEARER_TOKEN")
)
```
- `VALID_TOKEN` reads from `.env` file
- This is the token that backend will validate against

**Lines 363-369 (in main function):**
```go
if token := os.Getenv("BEARER_TOKEN"); token != "" {
    VALID_TOKEN = token
}
```
- Double-checks and loads token from environment

---

### Part 2: Expose Token via /config Endpoint

**Lines 352-358:**
```go
func configHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := map[string]string{
		"bearerToken": VALID_TOKEN,
	}
	respondWithJSON(w, http.StatusOK, response)
}
```
- Frontend calls this endpoint to get the token
- Returns: `{"bearerToken": "secret_token_12345"}`

**Line 379 (in main function):**
```go
http.HandleFunc("/config", configHandler)
```
- Registers the `/config` endpoint (no authentication required)

---

### Part 3: Validate Token (Authentication Middleware)

**Lines 85-109:**
```go
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Apply CORS first
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		authHeader := r.Header.Get("Authorization")
		
		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}
		
		// Extract token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		
		// Validate token
		if token != VALID_TOKEN {
			respondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}
		
		// Token is valid, proceed
		next(w, r)
	}
}
```

**Key Lines:**
- **Line 97:** `authHeader := r.Header.Get("Authorization")` - Gets the Authorization header
- **Line 99-102:** Checks if header is missing → Returns 401 error
- **Line 105:** `token := strings.TrimPrefix(authHeader, "Bearer ")` - Removes "Bearer " prefix
- **Line 108-111:** Compares token with `VALID_TOKEN` → If wrong, returns 401 error
- **Line 114:** If token is valid, allows request to proceed

---

### Part 4: Protected Routes (Require Token)

**Lines 382-407:**
```go
// User routes (with auth)
http.HandleFunc("/users", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        createUserHandler(w, r)
    } else {
        respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
    }
}))

http.HandleFunc("/users/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        getUserHandler(w, r)
    case http.MethodPut:
        updateUserHandler(w, r)
    case http.MethodPatch:
        patchUserHandler(w, r)
    case http.MethodDelete:
        deleteUserHandler(w, r)
    default:
        respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
    }
}))
```
- `authMiddleware(...)` wraps each handler
- Token is checked BEFORE the handler runs
- If token is invalid, request is rejected before reaching the handler

---

## 📁 File 3: `frontend/index.html` (Frontend - HTML/JavaScript)

### Part 1: Load Token from Backend

**Lines 460-463:**
```javascript
const API_URL = 'http://localhost:8080';
let TOKEN = '';
```
- `TOKEN` variable stores the Bearer token (initially empty)

**Lines 465-468:**
```javascript
window.addEventListener('load', async () => {
    await loadTokenFromBackend();
    checkServerStatus();
});
```
- When page loads, calls `loadTokenFromBackend()` function

**Lines 470-482:**
```javascript
async function loadTokenFromBackend() {
    try {
        const response = await fetch(`${API_URL}/config`);
        if (response.ok) {
            const data = await response.json();
            TOKEN = data.bearerToken;
        } else {
            console.error('❌ Failed to load token from backend');
        }
    } catch (error) {
        console.error('❌ Error loading token:', error);
    }
}
```
- **Line 472:** Calls `GET /config` endpoint
- **Line 475:** Extracts token from response: `data.bearerToken`
- **Line 476:** Stores token in `TOKEN` variable

---

### Part 2: Send Token in API Requests

**Example 1: GET Request (Lines 500-510)**
```javascript
async function getUser() {
    const userId = document.getElementById('userId').value;
    try {
        const response = await fetch(`${API_URL}/users/${userId}`, {
            headers: {
                'Authorization': `Bearer ${TOKEN}`
            }
        });
        const data = await response.json();
        displayResponse('userResult', data, !response.ok);
    } catch (error) {
        displayResponse('userResult', { error: error.message }, true);
    }
}
```
- **Line 504:** Sends request to `/users/{id}`
- **Line 505-507:** Includes `Authorization: Bearer {TOKEN}` header
- Backend receives this header and validates it

**Example 2: POST Request (Lines 514-537)**
```javascript
async function createUser() {
    const name = document.getElementById('createName').value;
    const email = document.getElementById('createEmail').value;
    const username = document.getElementById('createUsername').value;

    if (!name || !email || !username) {
        displayResponse('createResult', { error: 'All fields are required' }, true);
        return;
    }

    try {
        const response = await fetch(`${API_URL}/users`, {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${TOKEN}`,
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ name, email, username })
        });
        const data = await response.json();
        displayResponse('createResult', data, !response.ok);
        
        if (response.ok) {
            document.getElementById('createName').value = '';
            document.getElementById('createEmail').value = '';
            document.getElementById('createUsername').value = '';
        }
    } catch (error) {
        displayResponse('createResult', { error: error.message }, true);
    }
}
```
- **Line 525:** POST request to `/users`
- **Line 527:** Includes `Authorization: Bearer {TOKEN}` header
- **Line 528:** Also includes `Content-Type: application/json`

---

### Part 3: Token Verification Tests

**Lines 495-532 (Test WITH Token):**
```javascript
async function testWithToken() {
    try {
        const response = await fetch(`${API_URL}/users/1`, {
            headers: {
                'Authorization': `Bearer ${TOKEN}`
            }
        });
        const data = await response.json();
        
        if (response.ok) {
            displayResponse('tokenTestResult', {
                status: '✅ SUCCESS',
                message: 'Bearer token is working correctly!',
                token: TOKEN,
                response: data
            });
        } else {
            displayResponse('tokenTestResult', {
                status: '❌ FAILED',
                message: 'Token authentication failed',
                error: data
            }, true);
        }
    } catch (error) {
        displayResponse('tokenTestResult', { error: error.message }, true);
    }
}
```
- **Line 497-500:** Sends request WITH token
- **Line 504-509:** If successful (200 OK), shows success message
- **Line 510-515:** If failed, shows error

**Lines 534-560 (Test WITHOUT Token):**
```javascript
async function testWithoutToken() {
    try {
        const response = await fetch(`${API_URL}/users/1`);
        const data = await response.json();
        
        if (response.status === 401) {
            displayResponse('tokenTestResult', {
                status: '✅ CORRECT BEHAVIOR',
                message: 'Server correctly rejected request without token',
                httpStatus: response.status,
                error: data.error
            });
        } else {
            displayResponse('tokenTestResult', {
                status: '⚠️ WARNING',
                message: 'Server should have rejected this request!',
                httpStatus: response.status,
                response: data
            }, true);
        }
    } catch (error) {
        displayResponse('tokenTestResult', { error: error.message }, true);
    }
}
```
- **Line 536:** Sends request WITHOUT token (no Authorization header)
- **Line 539:** Checks if server returned 401 Unauthorized
- **Line 540-545:** If 401, shows "correct behavior" message
- **Line 546-551:** If NOT 401, shows warning (security issue!)

---

## 🔄 Complete Flow

### 1. Token Storage
```
.env (Line 3) → BEARER_TOKEN=secret_token_12345
```

### 2. Backend Loads Token
```
server.go (Line 13) → VALID_TOKEN = os.Getenv("BEARER_TOKEN")
```

### 3. Frontend Fetches Token
```
frontend/index.html (Line 472) → fetch('/config')
frontend/index.html (Line 476) → TOKEN = data.bearerToken
```

### 4. Frontend Sends Request with Token
```
frontend/index.html (Line 506) → 'Authorization': `Bearer ${TOKEN}`
```

### 5. Backend Validates Token
```
server.go (Line 97) → authHeader := r.Header.Get("Authorization")
server.go (Line 105) → token := strings.TrimPrefix(authHeader, "Bearer ")
server.go (Line 108) → if token != VALID_TOKEN { return 401 }
```

### 6. If Valid → Process Request
```
server.go (Line 114) → next(w, r)
```

### 7. If Invalid → Return Error
```
server.go (Line 109) → respondWithError(w, 401, "Invalid token")
```

---

## 🧪 How to Test

### Test 1: Check Token in Browser Console
1. Open `http://localhost:3000`
2. Press `F12` → Console tab
3. Type: `TOKEN` and press Enter
4. Should show: `"secret_token_12345"`

### Test 2: Check Token in Network Tab
1. Press `F12` → Network tab
2. Click "Get User" button
3. Click on the `users/1` request
4. Look at **Request Headers**
5. Should see: `Authorization: Bearer secret_token_12345`

### Test 3: Use Verification Buttons
1. Click "✅ Test WITH Token" → Should succeed
2. Click "❌ Test WITHOUT Token" → Should fail with 401

### Test 4: Test with curl
```bash
# WITH token (should work)
curl -H "Authorization: Bearer secret_token_12345" http://localhost:8080/users/1

# WITHOUT token (should fail)
curl http://localhost:8080/users/1
```

---

## 📝 Summary

| File | Lines | Purpose |
|------|-------|---------|
| `.env` | 3 | Store token |
| `server.go` | 13, 363-369 | Load token from environment |
| `server.go` | 352-358 | Expose token via `/config` |
| `server.go` | 85-114 | Validate token (middleware) |
| `server.go` | 382-407 | Protect routes with token |
| `index.html` | 470-482 | Fetch token from backend |
| `index.html` | 506, 527 | Send token in requests |
| `index.html` | 495-560 | Test token validation |
