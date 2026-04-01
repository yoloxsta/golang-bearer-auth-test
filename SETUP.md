# Setup Guide

## Prerequisites

- Go 1.21 or higher
- Git (for cloning)
## Installation

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/go-rest-api-lab.git
cd go-rest-api-lab
```

### 2. Create environment file

```bash
copy .env.example .env
```

The default `.env.example` is configured for local testing.

### 3. Install dependencies

```bash
go mod tidy
```

## Running the Project

### Method 1: Local Server Testing (Recommended)

**Terminal 1 - Start the API server:**
```bash
cd server
go run server.go
```

You should see:
```
🚀 REST API Server Started
Server running on: http://localhost:8080
Valid Bearer Token: secret_token_12345
```

**Terminal 2 - Run the client:**
```bash
go run .
```

Expected output:
```
=== Fetching User ===
User ID: 1
Name: John Doe
Email: john.doe@example.com
Username: johndoe

=== Fetching Post ===
Post ID: 1
User ID: 1
Title: My First Post
Body: This is the content of my first post...

✓ All requests completed successfully
```

### Method 2: Demo Mode (No Network Required)

```bash
go run . demo
```

This runs with mock data to demonstrate the code structure.

### Method 3: External API

Edit `.env` with your API credentials:

```env
API_BASE_URL=https://api.example.com
BEARER_TOKEN=your_actual_token
```

Then run:
```bash
go run .
```

## Testing with Postman

1. Start the local server (see Method 1 above)
2. Open Postman
3. Create a new GET request
4. URL: `http://localhost:8080/users/1`
5. Go to Authorization tab
6. Type: Bearer Token
7. Token: `secret_token_12345`
8. Click Send

See `server/README.md` for more testing examples.

## Project Structure

```
go-rest-api-lab/
├── main.go           # Client application entry point
├── demo.go           # Offline demo mode
├── go.mod            # Go module file
├── .env.example      # Environment template
├── LICENSE           # MIT License
├── README.md         # Main documentation
├── SETUP.md          # This file
├── client/           # HTTP client package
│   └── client.go     # Bearer token authentication
├── config/           # Configuration package
│   └── config.go     # .env file parser
├── models/           # Data models
│   └── models.go     # API response structures
└── server/           # Local test server
    ├── server.go     # REST API server
    └── README.md     # Server documentation
```

## Troubleshooting

### "Failed to open .env file"
- Make sure you copied `.env.example` to `.env`
- Check you're in the project root directory

### "Connection refused" or "Timeout"
- Make sure the server is running (Terminal 1)
- Check the server is on port 8080
- Try demo mode: `go run . demo`

### "Invalid token" (401 Unauthorized)
- Check your `.env` file has the correct token
- For local server, use: `secret_token_12345`

## Next Steps

- Modify `models/models.go` to match your API responses
- Update endpoints in `main.go`
- Add more HTTP methods (POST, PUT, DELETE) to `client/client.go`
- Add unit tests
- Deploy the server

## Support

For issues or questions, please open an issue on GitHub.
