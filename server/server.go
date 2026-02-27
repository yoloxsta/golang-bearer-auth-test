package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	PORT        = "8080"
	VALID_TOKEN = os.Getenv("BEARER_TOKEN")
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PatchUserRequest struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

type CreatePostRequest struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// CORS middleware to allow browser requests
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next(w, r)
	}
}

// Middleware to check Bearer token
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
		
		// Check if it starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			respondWithError(w, http.StatusUnauthorized, "Invalid Authorization format. Use: Bearer <token>")
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

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL path
	// For /users/1, we get ID 1
	// This is a simple implementation - in production use a router like gorilla/mux
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 3 {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	user, err := getUserByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	
	respondWithJSON(w, http.StatusOK, user)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	post, err := getPostByID(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Post not found")
		return
	}
	
	respondWithJSON(w, http.StatusOK, post)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Username == "" {
		respondWithError(w, http.StatusBadRequest, "Name, email, and username are required")
		return
	}
	
	// Create user in database
	user, err := createUser(req.Name, req.Email, req.Username)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			respondWithError(w, http.StatusConflict, "User with this email or username already exists")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to create user")
		}
		return
	}
	
	respondWithJSON(w, http.StatusCreated, user)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Validate required fields
	if req.Name == "" || req.Email == "" || req.Username == "" {
		respondWithError(w, http.StatusBadRequest, "Name, email, and username are required")
		return
	}
	
	// Update user in database
	user, err := updateUser(id, req.Name, req.Email, req.Username)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to update user")
		}
		return
	}
	
	respondWithJSON(w, http.StatusOK, user)
}

func patchUserHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	var req PatchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Update user in database
	user, err := patchUserDB(id, req.Name, req.Email)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to update user")
		}
		return
	}
	
	respondWithJSON(w, http.StatusOK, user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	err := deleteUser(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "User not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete user")
		}
		return
	}
	
	response := SuccessResponse{
		Message: "User deleted successfully",
		Data: map[string]int{
			"id": id,
		},
	}
	
	respondWithJSON(w, http.StatusOK, response)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Validate required fields
	if req.Title == "" || req.Body == "" {
		respondWithError(w, http.StatusBadRequest, "Title and body are required")
		return
	}
	
	// Create post in database
	post, err := createPost(req.UserID, req.Title, req.Body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to create post")
		return
	}
	
	respondWithJSON(w, http.StatusCreated, post)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Update post in database
	post, err := updatePost(id, req.UserID, req.Title, req.Body)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Post not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to update post")
		}
		return
	}
	
	respondWithJSON(w, http.StatusOK, post)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	var id int
	fmt.Sscanf(pathParts[2], "%d", &id)
	
	err := deletePost(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			respondWithError(w, http.StatusNotFound, "Post not found")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Failed to delete post")
		}
		return
	}
	
	response := SuccessResponse{
		Message: "Post deleted successfully",
		Data: map[string]int{
			"id": id,
		},
	}
	
	respondWithJSON(w, http.StatusOK, response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := map[string]string{
		"status": "ok",
		"message": "Server is running",
	}
	respondWithJSON(w, http.StatusOK, response)
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	response := map[string]string{
		"bearerToken": VALID_TOKEN,
	}
	respondWithJSON(w, http.StatusOK, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

func main() {
	// Load environment variables
	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}
	if token := os.Getenv("BEARER_TOKEN"); token != "" {
		VALID_TOKEN = token
	}
	
	// Initialize database
	if err := InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer CloseDB()
	
	// Routes - Health check (no auth)
	http.HandleFunc("/health", healthHandler)
	
	// Config endpoint (no auth) - returns bearer token for frontend
	http.HandleFunc("/config", configHandler)
	
	// User routes (with auth)
	// Handle /users (no trailing slash) for POST
	http.HandleFunc("/users", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createUserHandler(w, r)
		} else {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}))
	
	// Handle /users/{id} (with trailing slash)
	http.HandleFunc("/users/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Handle /users/{id}
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
	
	// Post routes (with auth)
	// Handle /posts (no trailing slash) for POST
	http.HandleFunc("/posts", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createPostHandler(w, r)
		} else {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}))
	
	// Handle /posts/{id} (with trailing slash)
	http.HandleFunc("/posts/", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		// Handle /posts/{id}
		switch r.Method {
		case http.MethodGet:
			getPostHandler(w, r)
		case http.MethodPut:
			updatePostHandler(w, r)
		case http.MethodDelete:
			deletePostHandler(w, r)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}))
	
	fmt.Println("========================================")
	fmt.Println("🚀 REST API Server Started")
	fmt.Println("========================================")
	fmt.Printf("Server running on: http://localhost:%s\n", PORT)
	fmt.Printf("Valid Bearer Token: %s\n", VALID_TOKEN)
	fmt.Println("\n📋 Available Endpoints:")
	fmt.Println("\n  Health:")
	fmt.Println("    GET    /health              - No auth required")
	fmt.Println("\n  Users:")
	fmt.Println("    GET    /users/{id}          - Get user by ID")
	fmt.Println("    POST   /users               - Create user")
	fmt.Println("    PUT    /users/{id}          - Update user (full)")
	fmt.Println("    PATCH  /users/{id}          - Update user (partial)")
	fmt.Println("    DELETE /users/{id}          - Delete user")
	fmt.Println("\n  Posts:")
	fmt.Println("    GET    /posts/{id}          - Get post by ID")
	fmt.Println("    POST   /posts               - Create post")
	fmt.Println("    PUT    /posts/{id}          - Update post")
	fmt.Println("    DELETE /posts/{id}          - Delete post")
	fmt.Println("\n🔐 All endpoints (except /health) require:")
	fmt.Println("    Authorization: Bearer secret_token_12345")
	fmt.Println("========================================")
	
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
