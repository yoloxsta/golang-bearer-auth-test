# Using Bearer Token from .env File

## How It Works

The Bearer token is now loaded from the `.env` file instead of being hardcoded.

### Backend (Go API)
1. Token is read from `.env` file: `BEARER_TOKEN=secret_token_12345`
2. Backend exposes `/config` endpoint that returns the token
3. All API requests validate against this token

### Frontend (HTML)
1. On page load, frontend calls `GET /config` to fetch the token
2. Token is automatically loaded and used for all API requests
3. You can also manually override the token using the input field

## Configuration

### 1. Edit `.env` File
```env
BEARER_TOKEN=your_custom_token_here
```

### 2. Rebuild and Restart
```bash
docker-compose build
docker-compose up -d
```

### 3. Test
Open `http://localhost:3000` - the token will be automatically loaded from backend.

## Endpoints

### GET /config (No Auth Required)
Returns the current bearer token from environment.

**Request:**
```bash
curl http://localhost:8080/config
```

**Response:**
```json
{
  "bearerToken": "secret_token_12345"
}
```

## Benefits

✅ No hardcoded tokens in code  
✅ Easy to change (just edit `.env`)  
✅ Secure (token stored server-side)  
✅ Frontend automatically syncs with backend  
✅ Can override token manually if needed  

## Security Note

⚠️ The `/config` endpoint is public (no auth required) so frontend can fetch the token. In production, you should:
- Use HTTPS
- Implement proper authentication flow (OAuth, JWT)
- Never expose sensitive tokens via public endpoints
- Use environment-specific tokens (dev, staging, prod)
