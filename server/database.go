package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "apiuser")
	password := getEnv("DB_PASSWORD", "apipassword")
	dbname := getEnv("DB_NAME", "restapi")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("✅ Database connected successfully")
	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// User database operations

func getUserByID(id int) (*User, error) {
	user := &User{}
	query := `SELECT id, name, email, username, created_at, updated_at FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func createUser(name, email, username string) (*User, error) {
	user := &User{}
	query := `INSERT INTO users (name, email, username) VALUES ($1, $2, $3) 
	          RETURNING id, name, email, username, created_at, updated_at`
	err := db.QueryRow(query, name, email, username).Scan(
		&user.ID, &user.Name, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func updateUser(id int, name, email, username string) (*User, error) {
	query := `UPDATE users SET name = $1, email = $2, username = $3, updated_at = CURRENT_TIMESTAMP 
	          WHERE id = $4 RETURNING id, name, email, username, created_at, updated_at`
	user := &User{}
	err := db.QueryRow(query, name, email, username, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func patchUserDB(id int, name, email *string) (*User, error) {
	// First get current user
	user, err := getUserByID(id)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	if name != nil {
		user.Name = *name
	}
	if email != nil {
		user.Email = *email
	}

	query := `UPDATE users SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP 
	          WHERE id = $3 RETURNING id, name, email, username, created_at, updated_at`
	err = db.QueryRow(query, user.Name, user.Email, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func deleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// Post database operations

func getPostByID(id int) (*Post, error) {
	post := &Post{}
	query := `SELECT id, user_id, title, body, created_at, updated_at FROM posts WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	}
	if err != nil {
		return nil, err
	}
	return post, nil
}

func createPost(userID int, title, body string) (*Post, error) {
	post := &Post{}
	query := `INSERT INTO posts (user_id, title, body) VALUES ($1, $2, $3) 
	          RETURNING id, user_id, title, body, created_at, updated_at`
	err := db.QueryRow(query, userID, title, body).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func updatePost(id, userID int, title, body string) (*Post, error) {
	query := `UPDATE posts SET user_id = $1, title = $2, body = $3, updated_at = CURRENT_TIMESTAMP 
	          WHERE id = $4 RETURNING id, user_id, title, body, created_at, updated_at`
	post := &Post{}
	err := db.QueryRow(query, userID, title, body, id).Scan(
		&post.ID, &post.UserID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("post not found")
	}
	if err != nil {
		return nil, err
	}
	return post, nil
}

func deletePost(id int) error {
	query := `DELETE FROM posts WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("post not found")
	}
	return nil
}
