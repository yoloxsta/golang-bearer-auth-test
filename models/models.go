package models

// User represents a user entity
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// Post represents a blog post entity
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// CreateUserRequest for POST /users
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// UpdateUserRequest for PUT /users/:id
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// PatchUserRequest for PATCH /users/:id
type PatchUserRequest struct {
	Name *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// CreatePostRequest for POST /posts
type CreatePostRequest struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// SuccessResponse for generic success messages
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
