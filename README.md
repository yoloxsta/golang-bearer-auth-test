# Go REST API Lab

A complete Go project demonstrating REST API client-server communication with Bearer token authentication.

## 🚀 Features

- ✅ **All HTTP Methods**: GET, POST, PUT, PATCH, DELETE
- ✅ **Bearer Token Authentication**: Industry-standard auth
- ✅ **Local Test Server**: Full REST API server included
- ✅ **Automated Testing**: Client tests all endpoints
- ✅ **Demo Mode**: Offline testing without network
- ✅ **Clean Architecture**: Modular, production-ready code
- ✅ **Zero Dependencies**: Pure Go stdlib only
- ✅ **Context & Timeouts**: Proper request handling
- ✅ **Error Handling**: Comprehensive error management

## Project Structure

```
.
├── .env.example              # Environment template
├── .gitignore                # Git ignore rules
├── LICENSE                   # MIT License
├── README.md                 # Main documentation
├── SETUP.md                  # Setup guide
├── TESTING.md                # Complete testing guide
├── postman_collection.json   # Postman collection
├── go.mod                    # Go module definition
├── main.go                   # Client entry point
├── demo.go                   # Offline demo mode
├── config/
│   └── config.go       # Configuration loader
├── client/
│   └── client.go       # HTTP client with Bearer auth
├── models/
│   └── models.go       # Response data structures
└── server/
    ├── server.go       # Local REST API server
    └── README.md       # Server documentation
```

## 📋 API Endpoints

### Server Endpoints (localhost:8080)

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | /health | Health check | ❌ |
| GET | /users/1 | Get user | ✅ |
| POST | /users | Create user | ✅ |
| PUT | /users/1 | Update user (full) | ✅ |
| PATCH | /users/1 | Update user (partial) | ✅ |
| DELETE | /users/1 | Delete user | ✅ |
| GET | /posts/1 | Get post | ✅ |
| POST | /posts | Create post | ✅ |
| PUT | /posts/1 | Update post | ✅ |
| DELETE | /posts/1 | Delete post | ✅ |

Valid Token: `secret_token_12345`

## Quick Start

### Option 1: Test with Local Server (Recommended)

**Terminal 1 - Start the server:**
```bash
cd server
go run server.go
```

**Terminal 2 - Create .env and run client:**
```bash
# Copy example config
copy .env.example .env

# Run the client
go run .
```

The `.env.example` is already configured for localhost!

### Option 2: Demo Mode (No Network)

```bash
go run . demo
```

### Option 3: External API

Edit `.env` with your API credentials:
```env
API_BASE_URL=https://api.example.com
BEARER_TOKEN=your_token_here
```

Then run:
```bash
go run .
```

## How It Works

### Configuration (config/config.go)
- Reads `.env` file line by line
- Parses key=value pairs
- Validates required fields
- Returns structured config

### HTTP Client (client/client.go)
- Creates reusable HTTP client with timeout
- Adds Authorization header: `Bearer <token>`
- Handles context cancellation
- Parses JSON responses
- Provides clean error messages

### Models (models/models.go)
- Defines structs matching API responses
- Uses JSON tags for field mapping
- Easy to extend for different endpoints

### Main Application (main.go)
- Loads configuration
- Creates API client
- Makes requests with context
- Formats and prints results

## Testing

### Automated Testing (Recommended)

**Terminal 1 - Start the server:**
```bash
cd server
go run server.go
```

**Terminal 2 - Run client tests:**
```bash
go run .
```

The client will automatically test all HTTP methods (GET, POST, PUT, PATCH, DELETE) and display results.

### Test with Postman

1. **Import Collection**: Import `postman_collection.json` into Postman
2. **Start Server**: `cd server && go run server.go`
3. **Run Tests**: Execute requests from the collection

The collection includes:
- All CRUD operations for Users and Posts
- Authentication tests (valid/invalid tokens)
- Pre-configured with localhost:8080 and valid token

See `TESTING.md` for detailed testing guide.

## Using with External APIs

### GitHub API Example
```env
API_BASE_URL=https://api.github.com
BEARER_TOKEN=ghp_your_github_token
```

### JSONPlaceholder (Testing)
```env
API_BASE_URL=https://jsonplaceholder.typicode.com
BEARER_TOKEN=test_token
```

Update models and endpoints in `main.go` accordingly.

## Key Concepts Explained

### 1. Bearer Token Authentication
```go
req.Header.Set("Authorization", "Bearer "+c.bearerToken)
```
Standard way to authenticate API requests using tokens.

### 2. Context with Timeout
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
```
Ensures requests don't hang indefinitely.

### 3. Error Wrapping
```go
return fmt.Errorf("failed to create request: %w", err)
```
Provides context while preserving original error.

### 4. JSON Unmarshaling
```go
json.Unmarshal(body, result)
```
Converts JSON bytes into Go structs.

## Extending the Project

### Add POST requests:
```go
func (c *APIClient) Post(ctx context.Context, endpoint string, data, result interface{}) error {
    // Implementation here
}
```

### Add custom headers:
```go
req.Header.Set("X-Custom-Header", "value")
```

### Add retry logic:
```go
for i := 0; i < maxRetries; i++ {
    // Retry logic
}
```

## Error Handling

The client handles:
- Network errors
- Timeout errors
- Non-200 status codes
- JSON parsing errors
- Missing configuration

All errors are wrapped with context for debugging.

## Best Practices Used

1. **Separation of concerns** - Config, client, models in separate packages
2. **Context usage** - Proper timeout and cancellation
3. **Error wrapping** - Clear error messages with context
4. **Resource cleanup** - defer for closing response bodies
5. **Configuration validation** - Check required fields early
6. **Type safety** - Strongly typed structs for responses

## Common Issues

### "Failed to open .env file"
- Make sure `.env` exists in project root
- Copy from `.env.example`

### "API returned status 401"
- Check your Bearer token is valid
- Verify token format in API documentation

### "context deadline exceeded"
- Increase timeout duration
- Check network connectivity
- Verify API endpoint is accessible

## Next Steps

- Add unit tests for client package
- Implement request/response logging
- Add rate limiting
- Create mock server for testing
- Add support for other HTTP methods (POST, PUT, DELETE)
