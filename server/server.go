package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	PORT         = "8080"
	VALID_TOKEN  = "secret_token_12345"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
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

// Middleware to check Bearer token
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Username: "johndoe",
	}
	
	respondWithJSON(w, http.StatusOK, user)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	post := Post{
		UserID: 1,
		ID:     1,
		Title:  "My First Post",
		Body:   "This is the content of my first post. Testing Bearer token authentication!",
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
	
	// Create new user (simulated)
	newUser := User{
		ID:       2,
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
	}
	
	respondWithJSON(w, http.StatusCreated, newUser)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
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
	
	// Update user (simulated)
	updatedUser := User{
		ID:       1,
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
	}
	
	respondWithJSON(w, http.StatusOK, updatedUser)
}

func patchUserHandler(w http.ResponseWriter, r *http.Request) {
	var req PatchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Start with existing user
	user := User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Username: "johndoe",
	}
	
	// Apply partial updates
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	
	respondWithJSON(w, http.StatusOK, user)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	response := SuccessResponse{
		Message: "User deleted successfully",
		Data: map[string]int{
			"id": 1,
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
	
	// Create new post (simulated)
	newPost := Post{
		UserID: req.UserID,
		ID:     2,
		Title:  req.Title,
		Body:   req.Body,
	}
	
	respondWithJSON(w, http.StatusCreated, newPost)
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	var req CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	// Update post (simulated)
	updatedPost := Post{
		UserID: req.UserID,
		ID:     1,
		Title:  req.Title,
		Body:   req.Body,
	}
	
	respondWithJSON(w, http.StatusOK, updatedPost)
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	response := SuccessResponse{
		Message: "Post deleted successfully",
		Data: map[string]int{
			"id": 1,
		},
	}
	
	respondWithJSON(w, http.StatusOK, response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "ok",
		"message": "Server is running",
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
	// Routes - Health check (no auth)
	http.HandleFunc("/health", healthHandler)
	
	// User routes (with auth)
	http.HandleFunc("/users/1", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
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
	
	http.HandleFunc("/users", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createUserHandler(w, r)
		default:
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	}))
	
	// Post routes (with auth)
	http.HandleFunc("/posts/1", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
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
	
	http.HandleFunc("/posts", authMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createPostHandler(w, r)
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
	fmt.Println("    GET    /users/1             - Get user")
	fmt.Println("    POST   /users               - Create user")
	fmt.Println("    PUT    /users/1             - Update user (full)")
	fmt.Println("    PATCH  /users/1             - Update user (partial)")
	fmt.Println("    DELETE /users/1             - Delete user")
	fmt.Println("\n  Posts:")
	fmt.Println("    GET    /posts/1             - Get post")
	fmt.Println("    POST   /posts               - Create post")
	fmt.Println("    PUT    /posts/1             - Update post")
	fmt.Println("    DELETE /posts/1             - Delete post")
	fmt.Println("\n🔐 All endpoints (except /health) require:")
	fmt.Println("    Authorization: Bearer secret_token_12345")
	fmt.Println("========================================")
	
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
