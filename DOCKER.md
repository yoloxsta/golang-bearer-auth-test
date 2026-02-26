# 🐳 Docker Setup Guide

Complete Docker setup for the Go REST API Lab.

## 🚀 Quick Start (3 Commands)

### 1. Build Images
```bash
docker-compose build
```

### 2. Start Services
```bash
docker-compose up -d
```

### 3. Open in Browser
- **API:** http://localhost:8080
- **Frontend:** http://localhost:3000

That's it! 🎉

---

## 📦 What Gets Deployed

### Services

1. **API Server** (Port 8080)
   - Go REST API
   - All HTTP methods
   - Bearer token auth
   - Built from source

2. **Frontend** (Port 3000)
   - HTML + JavaScript UI
   - Served by Nginx
   - Beautiful interface

---

## 🎯 Using Makefile (Easier)

### Start Everything
```bash
make up
```

### Stop Everything
```bash
make down
```

### View Logs
```bash
make logs
```

### Restart Services
```bash
make restart
```

### Test API
```bash
make test
```

### Clean Everything
```bash
make clean
```

### See All Commands
```bash
make help
```

---

## 🔧 Manual Docker Commands

### Build
```bash
docker-compose build
```

### Start (detached)
```bash
docker-compose up -d
```

### Start (with logs)
```bash
docker-compose up
```

### Stop
```bash
docker-compose down
```

### View Logs
```bash
docker-compose logs -f
```

### View API Logs Only
```bash
docker-compose logs -f api
```

### View Frontend Logs Only
```bash
docker-compose logs -f frontend
```

### Check Status
```bash
docker-compose ps
```

### Restart
```bash
docker-compose restart
```

### Remove Everything
```bash
docker-compose down -v --rmi all
```

---

## 🧪 Testing After Deployment

### Test 1: Health Check
```bash
curl http://localhost:8080/health
```

Expected:
```json
{
  "status": "ok",
  "message": "Server is running"
}
```

### Test 2: Get User
```bash
curl -H "Authorization: Bearer secret_token_12345" \
     http://localhost:8080/users/1
```

Expected:
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "username": "johndoe"
}
```

### Test 3: Frontend UI
Open browser: http://localhost:3000

You should see the beautiful purple interface!

---

## 📊 Architecture

```
┌─────────────────────────────────────┐
│         Docker Network              │
│                                     │
│  ┌──────────────┐  ┌─────────────┐ │
│  │   Frontend   │  │  API Server │ │
│  │   (Nginx)    │  │    (Go)     │ │
│  │   Port 3000  │  │  Port 8080  │ │
│  └──────────────┘  └─────────────┘ │
│         │                  │        │
└─────────┼──────────────────┼────────┘
          │                  │
          ▼                  ▼
    http://localhost:3000    http://localhost:8080
```

---

## 🔍 Container Details

### API Container
- **Name:** go-rest-api
- **Port:** 8080
- **Image:** Built from Dockerfile
- **Base:** golang:1.21-alpine
- **Size:** ~50MB

### Frontend Container
- **Name:** go-rest-frontend
- **Port:** 3000
- **Image:** nginx:alpine
- **Size:** ~25MB

---

## 🛠️ Configuration

### Environment Variables

Edit `docker-compose.yml`:

```yaml
environment:
  - PORT=8080
  - BEARER_TOKEN=your_token_here
```

### Change Ports

Edit `docker-compose.yml`:

```yaml
ports:
  - "9090:8080"  # API on port 9090
  - "4000:80"    # Frontend on port 4000
```

---

## 🐛 Troubleshooting

### Port Already in Use

**Problem:** Port 8080 or 3000 already in use

**Solution:**
```bash
# Stop existing services
docker-compose down

# Or change ports in docker-compose.yml
```

### Container Won't Start

**Problem:** Container exits immediately

**Solution:**
```bash
# Check logs
docker-compose logs api

# Rebuild
docker-compose build --no-cache
docker-compose up
```

### Can't Connect to API

**Problem:** Frontend can't reach API

**Solution:**
- Check both containers are running: `docker-compose ps`
- Check API logs: `docker-compose logs api`
- Verify network: `docker network ls`

### Permission Denied

**Problem:** Docker permission errors

**Solution (Windows):**
- Run Docker Desktop as Administrator
- Or add user to docker-users group

---

## 📝 Development Workflow

### 1. Make Code Changes
Edit your Go or HTML files

### 2. Rebuild
```bash
docker-compose build
```

### 3. Restart
```bash
docker-compose up -d
```

### 4. Test
```bash
make test
```

---

## 🚀 Production Deployment

### Build for Production
```bash
docker-compose -f docker-compose.yml build
```

### Push to Registry
```bash
# Tag image
docker tag go-rest-api:latest your-registry/go-rest-api:latest

# Push
docker push your-registry/go-rest-api:latest
```

### Deploy to Server
```bash
# On production server
docker-compose pull
docker-compose up -d
```

---

## 📊 Resource Usage

### Check Resource Usage
```bash
docker stats
```

### Typical Usage
- **API Container:** ~10MB RAM, <1% CPU
- **Frontend Container:** ~5MB RAM, <1% CPU

Very lightweight! ✅

---

## 🔐 Security Notes

### Production Checklist
- [ ] Change default Bearer token
- [ ] Use environment variables for secrets
- [ ] Enable HTTPS
- [ ] Add rate limiting
- [ ] Use Docker secrets
- [ ] Scan images for vulnerabilities

### Scan for Vulnerabilities
```bash
docker scan go-rest-api
```

---

## 📚 Docker Files

```
.
├── Dockerfile              # API server image
├── docker-compose.yml      # Service orchestration
├── .dockerignore          # Exclude files
├── Makefile               # Easy commands
└── DOCKER.md              # This file
```

---

## ✅ Verification Checklist

After running `docker-compose up -d`:

- [ ] Check containers running: `docker-compose ps`
- [ ] API responds: `curl http://localhost:8080/health`
- [ ] Frontend loads: Open http://localhost:3000
- [ ] Can create user via UI
- [ ] Can get user via UI
- [ ] All HTTP methods work
- [ ] Logs are clean: `docker-compose logs`

---

## 🎓 Docker Commands Reference

| Command | Description |
|---------|-------------|
| `docker-compose build` | Build images |
| `docker-compose up` | Start services |
| `docker-compose up -d` | Start in background |
| `docker-compose down` | Stop services |
| `docker-compose ps` | List containers |
| `docker-compose logs` | View logs |
| `docker-compose restart` | Restart services |
| `docker-compose exec api sh` | Shell into API |
| `docker-compose pull` | Pull latest images |

---

## 🎉 Benefits of Docker

✅ **Consistent Environment** - Works everywhere  
✅ **Easy Deployment** - One command to start  
✅ **Isolated** - No conflicts with other apps  
✅ **Portable** - Share with anyone  
✅ **Scalable** - Easy to add more containers  
✅ **Production-Ready** - Same setup everywhere  

---

## 🚀 Next Steps

1. ✅ Run `make up`
2. ✅ Open http://localhost:3000
3. ✅ Test the API
4. ✅ Push to Docker Hub
5. ✅ Deploy to production

---

## 📞 Quick Reference

```bash
# Start everything
make up

# Stop everything
make down

# View logs
make logs

# Test API
make test

# Clean everything
make clean
```

**API:** http://localhost:8080  
**Frontend:** http://localhost:3000  
**Token:** secret_token_12345

---

Happy Dockerizing! 🐳
