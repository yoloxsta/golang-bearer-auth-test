# 🗄️ PostgreSQL Database Integration

## ✅ What Changed

Your REST API now uses a **real PostgreSQL database** with full persistence!

### Before (Mock Data):
- ❌ Data was hardcoded
- ❌ Creating users didn't save them
- ❌ Only user ID 1 existed
- ❌ Data lost on restart

### After (Real Database):
- ✅ Data persists in PostgreSQL
- ✅ Create users and they're saved
- ✅ Get any user by ID (1, 2, 3, etc.)
- ✅ Data survives restarts
- ✅ Production-ready!

---

## 🐳 Containers Running

```
✅ go-rest-db (Port 5432) - PostgreSQL Database
✅ go-rest-api (Port 8080) - API Server
✅ go-rest-frontend (Port 3000) - Web UI
```

---

## 🎯 Test It Now!

### 1. Open Frontend
http://localhost:3000

### 2. Get Existing User
- Click "GET Get User" with ID 1
- You'll see John Doe (from sample data)

### 3. Create New User
- Fill in:
  - Name: Alice Johnson
  - Email: alice@example.com
  - Username: alice
- Click "POST Create User"
- You'll get ID 3 (or next available)

### 4. Get Your New User
- Change User ID to 3 (or the ID you got)
- Click "GET Get User"
- **You'll see your created user!** ✅

### 5. Update User
- Fill in new data
- Click "PUT Full Update"
- Data is updated in database!

### 6. Delete User
- Click "DELETE Delete User"
- User is removed from database!

---

## 📊 Database Schema

### Users Table
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Posts Table
```sql
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    title VARCHAR(500) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## 🌱 Sample Data

The database starts with:

**Users:**
- ID 1: John Doe (john.doe@example.com)
- ID 2: Jane Smith (jane.smith@example.com)

**Posts:**
- ID 1: "My First Post" (by John)
- ID 2: "Second Post" (by John)
- ID 3: "Jane's Post" (by Jane)

---

## 🔍 Database Operations

### Create User
```bash
curl -X POST \
  -H "Authorization: Bearer secret_token_12345" \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com","username":"alice"}' \
  http://localhost:8080/users
```

Response:
```json
{
  "id": 3,
  "name": "Alice",
  "email": "alice@example.com",
  "username": "alice",
  "created_at": "2026-02-25T03:11:00Z",
  "updated_at": "2026-02-25T03:11:00Z"
}
```

### Get User by ID
```bash
curl -H "Authorization: Bearer secret_token_12345" \
  http://localhost:8080/users/3
```

### Update User
```bash
curl -X PUT \
  -H "Authorization: Bearer secret_token_12345" \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Updated","email":"alice.new@example.com","username":"alice"}' \
  http://localhost:8080/users/3
```

### Delete User
```bash
curl -X DELETE \
  -H "Authorization: Bearer secret_token_12345" \
  http://localhost:8080/users/3
```

---

## 🗄️ Database Access

### Connect to PostgreSQL

```bash
# Using docker exec
docker exec -it go-rest-db psql -U apiuser -d restapi

# Or using psql locally
psql -h localhost -p 5432 -U apiuser -d restapi
```

Password: `apipassword`

### Useful SQL Commands

```sql
-- List all users
SELECT * FROM users;

-- List all posts
SELECT * FROM posts;

-- Count users
SELECT COUNT(*) FROM users;

-- Get user with posts
SELECT u.name, p.title 
FROM users u 
LEFT JOIN posts p ON u.id = p.user_id 
WHERE u.id = 1;

-- Delete all users (careful!)
DELETE FROM users;

-- Reset auto-increment
ALTER SEQUENCE users_id_seq RESTART WITH 1;
```

---

## 💾 Data Persistence

### Data is Stored in Docker Volume

```bash
# List volumes
docker volume ls

# Inspect volume
docker volume inspect go-token_postgres_data
```

### Data Survives:
- ✅ Container restart
- ✅ Docker restart
- ✅ Computer restart

### Data is Lost When:
- ❌ You run `docker-compose down -v` (removes volumes)
- ❌ You delete the volume manually

---

## 🔄 Reset Database

### Option 1: Restart Containers (keeps data)
```bash
docker-compose restart
```

### Option 2: Rebuild (keeps data)
```bash
docker-compose down
docker-compose up -d
```

### Option 3: Fresh Start (deletes data)
```bash
docker-compose down -v
docker-compose up -d
```

This will:
- Delete all data
- Recreate tables
- Insert sample data again

---

## 📁 New Files Added

```
✅ server/database.go      - Database operations
✅ database/init.sql       - Database schema & sample data
✅ go.mod                  - Added PostgreSQL driver
✅ docker-compose.yml      - Added PostgreSQL service
```

---

## 🎓 What You're Learning

### Database Concepts:
- SQL database integration
- CRUD operations (Create, Read, Update, Delete)
- Database connections
- Connection pooling
- Transactions
- Foreign keys
- Indexes

### Go Database:
- `database/sql` package
- PostgreSQL driver (`lib/pq`)
- Prepared statements
- Error handling
- Connection management

### Docker:
- Multi-container applications
- Service dependencies
- Health checks
- Volumes for persistence
- Environment variables

---

## 🚀 Production Features

Your API now has:

✅ **Real Database** - PostgreSQL  
✅ **Data Persistence** - Survives restarts  
✅ **Unique Constraints** - No duplicate emails/usernames  
✅ **Foreign Keys** - Posts linked to users  
✅ **Timestamps** - created_at, updated_at  
✅ **Auto-increment IDs** - SERIAL primary keys  
✅ **Indexes** - Fast lookups  
✅ **Triggers** - Auto-update timestamps  
✅ **Connection Pooling** - Efficient connections  
✅ **Health Checks** - Database monitoring  

---

## 🧪 Test Scenarios

### Scenario 1: Create and Retrieve
1. Create user "Bob" via POST
2. Note the ID (e.g., 3)
3. GET /users/3
4. See Bob's data! ✅

### Scenario 2: Update
1. GET /users/1 (John Doe)
2. PUT /users/1 with new name
3. GET /users/1 again
4. See updated name! ✅

### Scenario 3: Delete
1. Create user "Temp"
2. Note ID
3. DELETE /users/{id}
4. Try GET /users/{id}
5. Get 404 Not Found! ✅

### Scenario 4: Persistence
1. Create user "Persistent"
2. Stop containers: `docker-compose down`
3. Start again: `docker-compose up -d`
4. GET your user
5. Still there! ✅

---

## 🐛 Troubleshooting

### "User not found"
- Check the ID exists
- Run: `docker exec -it go-rest-db psql -U apiuser -d restapi -c "SELECT * FROM users;"`

### "Database connection failed"
- Check db container: `docker ps`
- Check logs: `docker-compose logs db`
- Wait for health check: Database takes ~5 seconds to start

### "Duplicate key error"
- Email or username already exists
- Use different email/username

### Can't connect to database
```bash
# Check database is healthy
docker ps

# Should show "healthy" status for go-rest-db
```

---

## 📊 Database Stats

```bash
# Connect to database
docker exec -it go-rest-db psql -U apiuser -d restapi

# Run queries
SELECT COUNT(*) FROM users;
SELECT COUNT(*) FROM posts;
SELECT pg_size_pretty(pg_database_size('restapi'));
```

---

## 🎉 Success!

Your REST API is now **production-ready** with:

- ✅ Real PostgreSQL database
- ✅ Full CRUD operations
- ✅ Data persistence
- ✅ Proper error handling
- ✅ Unique constraints
- ✅ Foreign key relationships
- ✅ Timestamps
- ✅ Connection pooling

**Test it now at http://localhost:3000!**

Create a user, get it by ID, update it, delete it - all data is real and persisted! 🚀

---

## 📚 Learn More

- PostgreSQL Docs: https://www.postgresql.org/docs/
- Go database/sql: https://pkg.go.dev/database/sql
- Docker Volumes: https://docs.docker.com/storage/volumes/
