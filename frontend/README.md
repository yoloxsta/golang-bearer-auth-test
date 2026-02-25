# Frontend UI for Go REST API

Beautiful, fully functional web interface to test your Go REST API.

## 🚀 Quick Start

### 1. Start the Go Server

```bash
cd server
go run server.go
```

Server must be running on `http://localhost:8080`

### 2. Open the Frontend

Simply open `index.html` in your browser:

**Option A: Double-click**
- Navigate to `frontend` folder
- Double-click `index.html`

**Option B: From terminal**
```bash
cd frontend
start index.html
```

**Option C: Drag and drop**
- Drag `index.html` into your browser

That's it! No build tools, no npm, no setup needed!

---

## ✨ Features

### User Management
- ✅ **GET User** - Retrieve user by ID
- ✅ **POST User** - Create new user
- ✅ **PUT User** - Full update (all fields)
- ✅ **PATCH User** - Partial update (only changed fields)
- ✅ **DELETE User** - Remove user

### Post Management
- ✅ **GET Post** - Retrieve post by ID
- ✅ **POST Post** - Create new post
- ✅ **PUT Post** - Update post
- ✅ **DELETE Post** - Remove post

### Additional Features
- ✅ Server status indicator (online/offline)
- ✅ Bearer token authentication (automatic)
- ✅ Beautiful, responsive UI
- ✅ JSON response viewer
- ✅ Error handling
- ✅ Form validation
- ✅ Confirmation dialogs for delete
- ✅ Color-coded HTTP methods
- ✅ Real-time API testing

---

## 🎨 UI Overview

### Header
- Shows server status (🟢 Online / 🔴 Offline)
- Automatically checks on page load

### Cards
Each card represents a different operation:
1. **User Management** - GET and DELETE user
2. **Create User** - POST new user
3. **Update User** - PUT (full) and PATCH (partial)
4. **Post Management** - GET and DELETE post
5. **Create Post** - POST new post
6. **Update Post** - PUT update

### Response Display
- Green text = Success
- Red text = Error
- JSON formatted for readability
- Scrollable for long responses

---

## 📖 How to Use

### Get User
1. Enter user ID (default: 1)
2. Click "GET Get User"
3. See user data in response box

### Create User
1. Fill in name, email, username
2. Click "POST Create User"
3. See created user with new ID
4. Form clears automatically on success

### Update User (Full)
1. Enter user ID to update
2. Fill in ALL fields (name, email, username)
3. Click "PUT Full Update"
4. See updated user data

### Update User (Partial)
1. Enter user ID to update
2. Fill in ONLY fields you want to change
3. Click "PATCH Partial Update"
4. See updated user (unchanged fields remain)

### Delete User
1. Enter user ID
2. Click "DELETE Delete User"
3. Confirm deletion
4. See success message

### Posts
Same pattern as users!

---

## 🔧 Technical Details

### API Configuration
```javascript
const API_URL = 'http://localhost:8080';
const TOKEN = 'secret_token_12345';
```

Change these if your server runs on different port or uses different token.

### HTTP Methods Used
- **GET** - Retrieve data
- **POST** - Create new resource
- **PUT** - Replace entire resource
- **PATCH** - Update partial resource
- **DELETE** - Remove resource

### Authentication
All requests (except health check) include:
```javascript
headers: {
    'Authorization': 'Bearer secret_token_12345'
}
```

### CORS
Server must have CORS enabled to allow browser requests.
Already configured in `server/server.go`!

---

## 🎯 Testing Scenarios

### Scenario 1: Complete User Flow
1. GET user 1 (see existing data)
2. POST new user (create Jane Smith)
3. PUT user 1 (full update)
4. PATCH user 1 (change only name)
5. DELETE user 1 (remove)

### Scenario 2: Error Handling
1. Try GET with invalid ID
2. Try POST with missing fields
3. Try DELETE non-existent user
4. See error messages in red

### Scenario 3: Posts
1. GET post 1
2. CREATE new post
3. UPDATE post 1
4. DELETE post 1

---

## 🐛 Troubleshooting

### "Server Offline" Status
**Problem:** Red status indicator  
**Solution:** 
- Make sure server is running: `cd server && go run server.go`
- Check server is on port 8080
- Refresh the page

### CORS Error in Console
**Problem:** "Access-Control-Allow-Origin" error  
**Solution:**
- Server has CORS enabled (already done!)
- If you modified server, check CORS headers

### No Response
**Problem:** Click button, nothing happens  
**Solution:**
- Open browser console (F12)
- Check for JavaScript errors
- Verify API_URL is correct

### 401 Unauthorized
**Problem:** All requests fail with 401  
**Solution:**
- Check TOKEN matches server token
- Default is `secret_token_12345`

---

## 💡 Customization

### Change Colors
Edit the CSS in `<style>` section:
```css
background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
```

### Change API URL
Edit JavaScript:
```javascript
const API_URL = 'http://your-server:port';
```

### Change Token
Edit JavaScript:
```javascript
const TOKEN = 'your_token_here';
```

### Add New Endpoints
1. Add new card in HTML
2. Add new function in JavaScript
3. Follow existing pattern

---

## 📱 Mobile Responsive

The UI is fully responsive:
- Desktop: 2-column grid
- Tablet: 2-column grid
- Mobile: 1-column stack

Test on any device!

---

## 🎓 Learning Points

This frontend demonstrates:
- ✅ Fetch API for HTTP requests
- ✅ Async/await for promises
- ✅ Bearer token authentication
- ✅ REST API methods (GET, POST, PUT, PATCH, DELETE)
- ✅ JSON parsing and display
- ✅ Error handling
- ✅ Form validation
- ✅ DOM manipulation
- ✅ CSS Grid layout
- ✅ Responsive design

---

## 🚀 Next Steps

### Enhancements You Could Add:
1. **List all users** - Show table of all users
2. **Search** - Filter users by name
3. **Pagination** - Handle many records
4. **Loading states** - Show spinner during requests
5. **Toast notifications** - Success/error popups
6. **Dark mode** - Toggle theme
7. **Request history** - Log all API calls
8. **Export data** - Download as JSON/CSV

---

## ✅ Checklist

Before using:
- [x] Server is running
- [x] CORS is enabled
- [x] index.html is opened in browser
- [x] Server status shows "Online"

Ready to test:
- [x] All HTTP methods work
- [x] Authentication works
- [x] Responses display correctly
- [x] Errors are handled
- [x] UI is responsive

---

## 🎉 You're Ready!

Open `index.html` in your browser and start testing your Go REST API with a beautiful UI!

No installation, no build process, just pure HTML + JavaScript + CSS.

Happy testing! 🚀
