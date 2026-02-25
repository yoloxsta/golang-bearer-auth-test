package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yourusername/go-rest-api-lab/client"
	"github.com/yourusername/go-rest-api-lab/config"
	"github.com/yourusername/go-rest-api-lab/models"
)

func main() {
	// Check if we should run in demo mode (no network required)
	if len(os.Args) > 1 && os.Args[1] == "demo" {
		RunOfflineDemo()
		return
	}

	// Load configuration from .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create API client with 10 second timeout
	apiClient := client.NewAPIClient(cfg.APIBaseURL, cfg.BearerToken, 10*time.Second)

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	fmt.Println("========================================")
	fmt.Println("🧪 REST API Client - Full Test Suite")
	fmt.Println("========================================\n")

	// GET requests
	fmt.Println("📥 GET REQUESTS")
	fmt.Println("----------------------------------------")
	testGetUser(ctx, apiClient)
	testGetPost(ctx, apiClient)

	// POST requests
	fmt.Println("\n📤 POST REQUESTS")
	fmt.Println("----------------------------------------")
	testCreateUser(ctx, apiClient)
	testCreatePost(ctx, apiClient)

	// PUT requests
	fmt.Println("\n🔄 PUT REQUESTS")
	fmt.Println("----------------------------------------")
	testUpdateUser(ctx, apiClient)
	testUpdatePost(ctx, apiClient)

	// PATCH requests
	fmt.Println("\n🔧 PATCH REQUESTS")
	fmt.Println("----------------------------------------")
	testPatchUser(ctx, apiClient)

	// DELETE requests
	fmt.Println("\n🗑️  DELETE REQUESTS")
	fmt.Println("----------------------------------------")
	testDeleteUser(ctx, apiClient)
	testDeletePost(ctx, apiClient)

	fmt.Println("\n========================================")
	fmt.Println("✅ All tests completed!")
	fmt.Println("========================================")
}

func testGetUser(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("GET /users/1")
	var user models.User
	
	if err := apiClient.Get(ctx, "/users/1", &user); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ User ID: %d\n", user.ID)
	fmt.Printf("  ✓ Name: %s\n", user.Name)
	fmt.Printf("  ✓ Email: %s\n", user.Email)
}

func testGetPost(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("\nGET /posts/1")
	var post models.Post
	
	if err := apiClient.Get(ctx, "/posts/1", &post); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ Post ID: %d\n", post.ID)
	fmt.Printf("  ✓ Title: %s\n", post.Title)
}

func testCreateUser(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("POST /users")
	
	newUser := models.CreateUserRequest{
		Name:     "Jane Smith",
		Email:    "jane.smith@example.com",
		Username: "janesmith",
	}
	
	var result models.User
	if err := apiClient.Post(ctx, "/users", newUser, &result); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ Created User ID: %d\n", result.ID)
	fmt.Printf("  ✓ Name: %s\n", result.Name)
	fmt.Printf("  ✓ Email: %s\n", result.Email)
}

func testCreatePost(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("\nPOST /posts")
	
	newPost := models.CreatePostRequest{
		UserID: 1,
		Title:  "New Post Title",
		Body:   "This is the body of the new post created via API",
	}
	
	var result models.Post
	if err := apiClient.Post(ctx, "/posts", newPost, &result); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ Created Post ID: %d\n", result.ID)
	fmt.Printf("  ✓ Title: %s\n", result.Title)
}

func testUpdateUser(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("PUT /users/1")
	
	updateData := models.UpdateUserRequest{
		Name:     "John Updated",
		Email:    "john.updated@example.com",
		Username: "johnupdated",
	}
	
	var result models.User
	if err := apiClient.Put(ctx, "/users/1", updateData, &result); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ Updated User ID: %d\n", result.ID)
	fmt.Printf("  ✓ New Name: %s\n", result.Name)
	fmt.Printf("  ✓ New Email: %s\n", result.Email)
}

func testUpdatePost(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("\nPUT /posts/1")
	
	updateData := models.CreatePostRequest{
		UserID: 1,
		Title:  "Updated Post Title",
		Body:   "This post has been updated via PUT request",
	}
	
	var result models.Post
	if err := apiClient.Put(ctx, "/posts/1", updateData, &result); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ Updated Post ID: %d\n", result.ID)
	fmt.Printf("  ✓ New Title: %s\n", result.Title)
}

func testPatchUser(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("PATCH /users/1")
	
	newName := "John Patched"
	patchData := models.PatchUserRequest{
		Name: &newName,
	}
	
	var result models.User
	if err := apiClient.Patch(ctx, "/users/1", patchData, &result); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("  ✓ Patched User ID: %d\n", result.ID)
	fmt.Printf("  ✓ New Name: %s\n", result.Name)
	fmt.Printf("  ✓ Email unchanged: %s\n", result.Email)
}

func testDeleteUser(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("DELETE /users/1")
	
	if err := apiClient.Delete(ctx, "/users/1"); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Println("  ✓ User deleted successfully")
}

func testDeletePost(ctx context.Context, apiClient *client.APIClient) {
	fmt.Println("\nDELETE /posts/1")
	
	if err := apiClient.Delete(ctx, "/posts/1"); err != nil {
		log.Printf("  ❌ Error: %v\n", err)
		return
	}

	fmt.Println("  ✓ Post deleted successfully")
}
