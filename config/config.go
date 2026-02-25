package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	APIBaseURL   string
	BearerToken  string
}

// LoadConfig reads configuration from .env file
func LoadConfig() (*Config, error) {
	config := &Config{}
	
	file, err := os.Open(".env")
	if err != nil {
		return nil, fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		
		switch key {
		case "API_BASE_URL":
			config.APIBaseURL = value
		case "BEARER_TOKEN":
			config.BearerToken = value
		}
	}
	
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %w", err)
	}
	
	// Validate required fields
	if config.APIBaseURL == "" {
		return nil, fmt.Errorf("API_BASE_URL is required")
	}
	if config.BearerToken == "" {
		return nil, fmt.Errorf("BEARER_TOKEN is required")
	}
	
	return config, nil
}
