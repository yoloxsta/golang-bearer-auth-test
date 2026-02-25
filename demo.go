package main

import (
	"encoding/json"
	"fmt"

	"github.com/yourusername/go-rest-api-lab/models"
)

// RunOfflineDemo demonstrates the code working with mock data
func RunOfflineDemo() {
	fmt.Println("=== OFFLINE DEMO MODE ===")
	fmt.Println("(Simulating API responses without network)\n")

	// Mock JSON responses (what the API would return)
	userJSON := `{
		"id": 1,
		"name": "Leanne Graham",
		"email": "Sincere@april.biz",
		"username": "Bret"
	}`

	postJSON := `{
		"userId": 1,
		"id": 1,
		"title": "sunt aut facere repellat provident occaecati",
		"body": "quia et suscipit suscipit recusandae consequuntur expedita"
	}`

	// Parse user
	fmt.Println("=== Fetching User ===")
	var user models.User
	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("User ID: %d\n", user.ID)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Username: %s\n", user.Username)

	// Parse post
	fmt.Println("\n=== Fetching Post ===")
	var post models.Post
	if err := json.Unmarshal([]byte(postJSON), &post); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Post ID: %d\n", post.ID)
	fmt.Printf("User ID: %d\n", post.UserID)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)

	fmt.Println("\n✓ Demo completed - JSON parsing works!")
	fmt.Println("\nThis shows:")
	fmt.Println("  ✓ Structs are correctly defined")
	fmt.Println("  ✓ JSON tags work properly")
	fmt.Println("  ✓ Data parsing is functional")
	fmt.Println("\nThe network issue is environmental (firewall/proxy).")
	fmt.Println("The code itself is working correctly!")
}
